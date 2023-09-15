from setuptools import setup, find_packages

setup(
    name='lunara-game-framework',
    version='0.1.0',
    description='Lunara Game Framework - A Python game development framework for Hypersonic Games',
    long_description='Lunara Game Framework is a Python game development framework for Hypersonic Games that is open sorce so that you can contribute too or use the framework that powers most of our games.',
    author='Hypersonic Games',
    url='https://github.com/A-Boring-Square/Lunara',
    license='GPL-3.0',
    packages=find_packages(),
    install_requires=[
        # Add your dependencies here
    ],
    classifiers=[
        'Development Status :: 3 - Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
        'Programming Language :: Python :: 3.8',
        'Programming Language :: Python :: 3.9',
    ],
    keywords='game framework lunara',
)
