package main

import (
    "log"
    "os"
    "os/signal"
    "sync"
    "syscall"

    "github.com/goav/av"
    "runtime"
)

// AudioPlayer represents the audio player instance.
type AudioPlayer struct {
    formatContext    *av.FormatContext
    audioStream      *av.Stream
    audioCodecContext *av.CodecContext
    audioCodec       *av.Codec
    audioFrameChan   chan *av.AudioFrame
    wg               sync.WaitGroup
    playing          bool
}

// NewAudioPlayer initializes a new AudioPlayer instance.
func NewAudioPlayer(inputFile string) (*AudioPlayer, error) {
    player := &AudioPlayer{}

    // Initialize FFmpeg.
    if err := av.Init(); err != nil {
        return nil, err
    }

    // Open the audio file.
    formatContext, err := av.OpenInputFile(inputFile)
    if err != nil {
        return nil, err
    }
    player.formatContext = formatContext

    // Find the audio stream.
    audioStream, err := formatContext.GetBestStream(av.AVMEDIA_TYPE_AUDIO)
    if err != nil {
        return nil, err
    }
    player.audioStream = audioStream

    // Initialize the audio decoder.
    audioCodecContext := audioStream.CodecContext()
    audioCodec, err := av.NewCodec(audioCodecContext.CodecID())
    if err != nil {
        return nil, err
    }
    if err := audioCodecContext.Open(audioCodec, nil); err != nil {
        return nil, err
    }
    player.audioCodecContext = audioCodecContext
    player.audioCodec = audioCodec

    // Create a channel for audio frames.
    player.audioFrameChan = make(chan *av.AudioFrame)

    // Start decoding audio packets.
    go player.decodeAudioPackets()

    return player, nil
}

// Play starts playing the audio.
func (player *AudioPlayer) Play() {
    if !player.playing {
        player.playing = true
        player.playAudioFrames()
    }
}

// Repeat plays the audio in a loop.
func (player *AudioPlayer) Repeat() {
    for {
        player.Play()
    }
}

// Info returns information about the audio stream.
func (player *AudioPlayer) Info() av.CodecContext {
    return *player.audioCodecContext
}

// Close releases resources and stops playback.
func (player *AudioPlayer) Close() {
    player.playing = false
    close(player.audioFrameChan)
    player.wg.Wait()
    player.audioCodecContext.Free()
    player.audioStream.Close()
    player.formatContext.Close()
    av.Exit()
}

// decodeAudioPackets decodes audio packets and sends audio frames to the channel.
func (player *AudioPlayer) decodeAudioPackets() {
    defer player.Close()
    for {
        packet, err := player.formatContext.ReadPacket()
        if err != nil {
            log.Printf("Error reading packet: %v", err)
            break
        }

        if packet.StreamIndex() == player.audioStream.Index() {
            audioFrame, err := player.audioCodecContext.Decode(packet)
            if err != nil {
                log.Printf("Error decoding audio frame: %v", err)
                continue
            }
            player.audioFrameChan <- audioFrame
        }

        packet.Free()
    }
}

// playAudioFrames plays audio frames.
func (player *AudioPlayer) playAudioFrames() {
    for i := 0; i < runtime.NumCPU(); i++ {
        player.wg.Add(1)
        go func() {
            defer player.wg.Done()
            for audioFrame := range player.audioFrameChan {
                // Implement audio playback logic here.
                // You may use audioFrame.Data to access audio data.
                // Example: playAudioData(audioFrame.Data)
            }
        }()
    }
}

func main() {
    // Initialize the audio player.
    inputFile := "input_audio.mp3"
    player, err := NewAudioPlayer(inputFile)
    if err != nil {
        log.Fatalf("Error initializing audio player: %v", err)
    }
    defer player.Close()

    // Play the audio.
    player.Play()

    // Handle cleanup and shutdown gracefully.
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
    <-c
    log.Println("Received termination signal. Cleaning up...")
}
