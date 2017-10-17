import discord
import random
from discord.ext import commands


'''For a list of built-in checks:
https://discordpy.readthedocs.io/en/rewrite/ext/commands/api.html#checks
You could also create your own custom checks. Check out:
https://github.com/Rapptz/discord.py/blob/master/discord/ext/commands/core.py#L689
For a list of events:
http://discordpy.readthedocs.io/en/rewrite/api.html#event-reference
http://discordpy.readthedocs.io/en/rewrite/ext/commands/api.html#event-reference
'''


class RNG:
	'''RNG'''
	def __init__(self, bot):
		self.bot = bot

	@commands.command(name='flip') 
	async def flip(self, ctx):
		'''Flip a coin.'''
		await ctx.send('`{}`'.format(random.choice(("Heads", "Tails"))))
	
	@commands.command(name='choice', aliases=['choose'])
	async def choose(self, ctx, *choices: str):
		'''Pick from multiple choices'''
		await ctx.send("`{}`".format(random.choice(choices)))

def setup(bot):
	bot.add_cog(RNG(bot))