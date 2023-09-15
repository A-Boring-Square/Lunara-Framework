import pygame as pg
import moderngl as mgl
import sys

class GraphicsEngine:
    def __init__(self, win_size=(500, 500)):
        pg.init()
        self.WIN_SIZE = win_size
        pg.display.gl_set_attribute(pg.GL_CONTEXT_MAJOR_VERSION, 3)
        pg.display.gl_set_attribute(pg.GL_CONTEXT_MINOR_VERSION, 3)
        pg.display.gl_set_attribute(pg.GL_CONTEXT_PROFILE_MASK, pg.GL_CONTEXT_PROFILE_CORE)
        pg.display.set_mode(self.WIN_SIZE, flags=pg.OPENGL | pg.DOUBLEBUF)
        self.ctx = mgl.create_context()
        self.clock = pg.time.Clock()
    
    def render(self):
        self.ctx.clear(color=(0.88, 0.16, 0.18))
        pg.display.flip()
    
    def run(self):
        running = True  # Variable to control the game loop
        while running:
            self.render()
            self.clock.tick(60)

            for event in pg.event.get():
                if event.type == pg.QUIT or (event.type == pg.KEYDOWN and event.key == pg.K_ESCAPE):
                    running = False  # Set running to False to exit the loop

        pg.quit()  # Quit Pygame properly
        sys.exit()

if __name__ == '__main__':
    app = GraphicsEngine()
    app.run()