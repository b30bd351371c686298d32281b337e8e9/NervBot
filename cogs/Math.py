import discord
import math
from discord.ext import commands



class Math:
	'''Math'''
	def __init__(self, bot):
		self.bot = bot
	
	@commands.command(name='math')
	async def meth(self):
		await self.bot.say("`Math is hard.` :cry:")
	
	@commands.command(name='add', aliases=['addition'])
	async def add(self, left : int, right : int):
		'''Adds two numbers together.'''
		await self.bot.say("`Difference: " + str(left - right) + "`")
	
	@commands.command(name='sub', aliases=['subtraction'])
	async def add(self, left : int, right : int):
		'''Adds two numbers together.'''
		await self.bot.say("`Summation: " + str(left + right) + "`")
	
	@commands.command(name='mult')
	async def mult(self, left : int, right: int):
		'''Multiples two numbers together.'''
		await self.bot.say("`Product: " + str(left * right) + "`")
	
	@commands.command(name='div', aliases=['divide'])
	async def div(self, left : int, right: int):
		'''Divides two numbers.'''
		await self.bot.say("`Quotients: " + str(left / right) + "`")
	
def setup(bot):
    bot.add_cog(Math(bot))