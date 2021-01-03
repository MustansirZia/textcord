# Generate Discord Token

This section deals with generating of what's called a Discord token. A token essentially is a long string of letters and is another way of gaining access to Discord besides the usual method of a username password pair. 
<br />
Although there are ways to generate a token for your own Discord account, for this excercise however, we would instead create a bot token or a token that belongs to a **Bot**. 
<br />
<br />
*Why should we need to use a bot account and not our own?*
> Well, it's generally bad practice as per Discord and may even lead to account deactivation if used excessively. Not to mention it could also compromise your account if you accidently expose it somewhere. 
<br />
<br />
With a bot account however, the element of risk is low as you can limit a bot's access to only a certain number of servers or channels. Discord also recommends using a bot account for any activity that's automated.


## Building the Bot
To get a Bot token we need to, well add a Bot, right?
1. The first step is log into to [Discord](https://discordapp.com) and head over to the [Applications](https://discordapp.com/developers/applications) page.
<br />

2. Tap on the **New Application** button to register a new application.
<br />

![New Discord Application](/discord-1.png)

3. In the dialog give it a name you like. If you're unsure just type *Simcord* and hit **Create**.
<br />

![Created Application](/discord-2.png)
*The newly created application would like this.*
<br />

4. Tap on the application tile you just created, find the **Bot** section in the left hand **Settings** pane and tap on it. 
<br />

![Bot](/discord-3.png)
*You should see a view like this.*
<br />

5. Tap on the **Add Bot** button and confirm your action in the dialog that follows.

6. You've just created a Bot! 🎉 Now is the time to get creative and choose a name for this Bot. 
<br /> *This is the name the people will in channels when you post a text using SMS so make sure this sounds dope. You can also change your bot's profile picture from here.* 

![Build a Bot](/discord-4.png)
*Your newly created Bot.*

<br /> Click on **Save Changes** and this should be all about creating a bot.

<br />

## Adding the Bot To Your Servers
Since our super dope bot is ready we can now add it to one or more Discord servers.

1. On the same screen, tap on the <b>OAuth2</b> section in the left hand **Settings** pane. 

2. Scroll down to the <b>Scopes</b> section and check the <b>"bot"</b>option. 

3. After checking the bot option, the <b>Bot Permissions</b> section will appear. In this section you need to enable a few checkboxes. This is shown in the image below.

![Bot Permissions](/discord-5.png)
*Bot permissions.*

4. Pass the link that's in bottom of the <b>bot</b> section to the admin of the discord server. Open the link yourself if you're yourself the admin or have <b>Manage Server</b> permission. 

![Bot URL](/discord-6.png)
*Bot URL.*

5. After opening the link a dialog will appear which will contain the list of servers that this can be added to. 

![Add to Server](/discord-7.png)
*Add to Server.*

6. Choose the server and tap on <b>Continue</b>. 

7. On the next screen confirm all permissions and tap on <b>Authorize</b>.

The bot is now added to your server 🎉!

## Getting the Token
The bot has now been added to your server but there is one additional step we still need to do. We need to get hold of the bot token.


1. In the left hand pane choose <b>Bot</b>. This is same section wherein you created the bot.

2. Beneath the Token, tap on the <b>Copy</b> button to copy the token to your clipboard. You can tap on the <b>Click to Reveal Token</b> to see the token yourself.

## Proceed Further?
::: tip Proceed Further?
You now have successfully created a bot, added the bot to your server(s) and as well taken hold of the our Super Dope Bot's token! 
<br />
We can now proceed further.
<br />
Head over to the [Go Live](/deploy/) page or click on the link at the bottom.
:::