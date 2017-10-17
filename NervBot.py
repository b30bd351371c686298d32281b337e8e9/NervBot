import discord
import asyncio
import websockets
import json
import sys, traceback
from time import gmtime, strftime
import Creds

from discord.ext.commands import Bot
from discord.ext import commands

#client = discord.Client()

bot_prefix = "^"
startup_extensions = ["cogs.Audio", "cogs.RNG", "cogs.Math"]
bot = commands.Bot(command_prefix=commands.when_mentioned_or(bot_prefix), description="2nd best Discord bot ever.")

def concat_result(result):
	return "`{}`".format(result)

@bot.event
async def on_ready():
	'''http://discordpy.readthedocs.io/en/rewrite/api.html#discord.on_ready'''

	print(f'Logged in as: {bot.user.name} with User ID {bot.user.id}\nVersion: {discord.__version__}\n')

	# Changes our bots Playing Status. type=1(streaming) for a standard game you could remove type and url.
	await bot.change_presence(game=discord.Game(name='FINAL FANTASY XVI', type=1, url='https://twitch.tv/LiquidData'))

  tabs-consistency
	# Here we load our extensions listed above in [startup_extensions].
	if __name__ == '__main__':
		for extension in startup_extensions:
			try:
				bot.load_extension(extension)
			except Exception as e:
				print(f'Failed to load extension {extension}.', file=sys.stderr)
				traceback.print_exc()
	print(f'Successfully logged in and booted...!')
 
    # Here we load our extensions listed above in [startup_extensions].
    for extension in startup_extensions:
        try:
            bot.load_extension(extension)
        except Exception as e:
            print(f'Failed to load extension {extension}.', file=sys.stderr)
            traceback.print_exc()
    print(f'Successfully logged in and booted...!')
    master

@bot.command()
async def load(extension_name : str):
	'''Loads an extension.'''
	try:
		bot.load_extension(extension_name)
	except (AttributeError, ImportError) as e:
		await bot.say("```py\n{}: {}\n```".format(type(e).__name__, str(e)))
		return
	await bot.say("{} loaded.".format(extension_name))

@bot.command()
async def unload(extension_name : str):
	'''Unloads an extension.'''
	bot.unload_extension(extension_name)
	await bot.say("{} unloaded.".format(extension_name))
	
@bot.command()
async def repeat(times : int, content='repeating...'):
	'''Repeats a message multiple times.'''
	for i in range(times):
		await bot.say(content)
		
@bot.command()
async def aimer():
	'''The only command you'll ever need.'''
	embed=discord.Embed(title="😍😍😍", description="***All things エメ***", url="https://line.me/ti/p/%40aimer", color=0x6666ff)
	embed.set_author(name="Aimer", url='http://www.aimer-web.jp/discography/', icon_url='https://pbs.twimg.com/profile_images/891898233684762624/uv3kAYy8_400x400.jpg')
	embed.set_thumbnail(url='https://i.imgur.com/P1GyRME.png')
	embed.add_field(name="Twitter Account", value="https://twitter.com/aimer_and_staff?lang=en", inline=True)
	embed.add_field(name="SMEJ YouTube Channel", value="https://www.youtube.com/user/aimerSMEJ", inline=False)
	embed.add_field(name="VEVO YouTube Channel", value="https://www.youtube.com/user/AimerOfficialVEVO", inline=True)
	embed.add_field(name="Reddit", value="https://www.reddit.com/r/Aimer/", inline=True)
	embed.add_field(name="Spotify", value="https://open.spotify.com/artist/0bAsR2unSRpn6BQPEnNlZm", inline=True)
	embed.add_field(name="Wikipedia", value="https://en.wikipedia.org/wiki/Aimer", inline=True)
	embed.add_field(name="Last.fm", value="https://www.last.fm/music/Aimer", inline=True)
	embed.add_field(name="Google+", value="https://plus.google.com/+aimer", inline=True)
	embed.set_footer(text="Created | {} GMT".format(strftime("%Y-%m-%d %H:%M:%S", gmtime())), icon_url="https://i.imgur.com/P1GyRME.png")
	await bot.say(embed=embed)

if __name__ == '__main__':
	bot.run(Creds.Token, bot=True, reconnect=True)