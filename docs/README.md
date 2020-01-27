---
home: true
heroImage: /discord-logo.svg
heroText: Simcord
tagline: Send or receive texts from Discord via SMS.
actionText: Get Started →
actionLink: /introduction/
features:
- title: What Does It Do?
  details: Simcord is tool that can be used for sending messages to and receiving messages from any Discord text channel via SMS.
- title: What Is The Use?
  details: It could allow you to access a discord chat without an active internet connection.
- title: What Does It Cost?
  details: We'll be using a generous free tier from Twilio. We'll also deploy and host it for free. You'd only pay your carrier for the SMSes you send.
footer: MIT Licensed | Copyright © 2020-present Mustansir Zia
---

## Proof Of Concept
Simcord is a serverless function that acts a webhook for incoming SMSes that come from your phone. It first parses the SMS body, then queries Discord using the Discord API and finally returns the response by sending a reply SMS to the number that sent the original SMS.

<br/> 

## What It Actually Is?
Simcord is a serverless function that acts a webhook for incoming SMSes that come from your phone. It first parses the SMS body, then queries Discord using the Discord [API](https://discordapp.com/developers/docs/intro) and finally returns the response by sending a reply to the phone number that sent the original SMS.

<br/> 

> Cellphone icon made by <a href="https://www.flaticon.com/authors/pixel-perfect">Pixel perfect</a> from <a href="http://www.flaticon.com">www.flaticon.com</a>.

