from pydub import AudioSegment
from pydub.playback import play

class AudioPlayer:
    def __init__(self, audio_file):
        self.audio = AudioSegment.from_file(audio_file)
        self.playing = False

    def play(self):
        if not self.playing:
            play(self.audio)
            self.playing = True

    def pause(self):
        if self.playing:
            play(AudioSegment.empty())
            self.playing = False

    def speed_up(self, factor=1.2):
        self.audio = self.audio.speedup(playback_speed=factor)
        self.play()

    def slow_down(self, factor=0.8):
        self.audio = self.audio.speedup(playback_speed=factor)
        self.play()

    def repeat(self, times=1):
        for _ in range(times):
            self.play()

if __name__ == "__main__":
    audio_file = "your_audio.mp3"  # Replace with your audio file path
    player = AudioPlayer(audio_file)

    print("Commands:")
    print("P - Play")
    print("S - Pause")
    print("+ - Speed Up")
    print("- - Slow Down")
    print("R - Repeat")
    print("Esc - Exit")

    while True:
        command = input("Enter a command: ").lower()
        if command == "p":
            player.play()
        elif command == "s":
            player.pause()
        elif command == "+":
            player.speed_up()
        elif command == "-":
            player.slow_down()
        elif command == "r":
            repeat_times = int(input("Enter number of times to repeat: "))
            player.repeat(repeat_times)
        elif command == "esc":
            break
