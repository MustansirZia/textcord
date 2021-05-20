# Go Live and Deploy

It's time to go live and deploy Textcord to the cloud! 
<br />

Why do we need to do this?
> Textcord under the hood is a serverless function that needs to live somewhere on the Internet so it can respond to requests and form the bridge between our inbox and Discord. This section will help us accomplish this. 

Unlike the rest of the steps this step is in fact the easiest and is also my favourite because it's incredibly simple. :)

## Make a Vercel Account

1. Login to your GitHub, GitLab or BitBucket account. Create an account if you don't have one from [here](https://github.com/join?source=login).

2. Now, head over [here](https://vercel.com/login) and create a Vercel account. You can use your existing existing GitHub, GitLab or BitBucket account to signup quickly.

## Deployment

1. Tap on this
[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/git/external?repository-url=https%3A%2F%2Fgithub.com%2FMustansirZia%2Fsimcord&env=DISCORD_TOKEN&envDescription=Add%20the%20discord%20token%20from%20the%20previous%20step.&envLink=https%3A%2F%2Fsimcord.now.sh%2Fdiscord%2F&project-name=my-simcord&repository-name=my-simcord)

2. Now follow along. Vercel will basically clone the main Textcord repo into your own GitHub, GitLab or BitBucket account and will use that for deployment. 

3. Before deploying, you'll probably see a screen like this. 

![Deploy Textcord](/deploy-1.png)
*Deploy Textcord.*

4. There's a <b>Required Environment Variables</b> section and after collapsing it you'll notice it houses just a single variable called `DISCORD_TOKEN`.

5. By now you've probably guessed it. This is the place where in we need to add the bot token we generated in the previous [Generate Discord Token](/discord) page. To add the bot token the value for this variable must be of the format `Bot <BOT_TOKEN>`. For example, your token is `12345678` then the value which you need to save is `Bot 12345678`. Please note if you're using token from your own Discord account then you need to omit the `Bot` part and use the token as is.

6. Tap on deploy and wait for the project to become live. Your own version of Textcord was successfully deployed 🎉. 

We have now deployed a serverless API without breaking a sweat!

![Successful Deployment](/deploy-2.png)
*Successful Deployment.*

7. Upon deployment you should see a page like this. Tap on <b>Visit</b> to see this very website hosted now from your own account! 

8. Note down the address for this website. We're going to need this later when we hook up our SMS gateway. It should be something like `https://my-textcord.vercel.app`.

## Ready for More?
::: tip Ready for More?
Phew! We have covered a lot so far. We have built a bot and also deployed our own version of Textcord on the cloud. 
<br />
Now, only the last milestone remains. 
<br />
Head over to the [SMS Gateway Configuration](/sms/) page or click on the link at the bottom.
:::
