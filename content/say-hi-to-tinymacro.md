+++
title = 'Say hi to Tinymacro'
date = 2024-02-01
draft = true
+++

Two weeks ago I shared a blog post where I introduced and idea that has been on my mind for some time. Here is the blog post: [Building a startup](https://denniscmartin.com/building-a-startup/). This time, I'd like to give you a brief introduction to this idea.

See, I really like finance. I remember working as a Crossfit trainer and spending all my free time in the gym reading articles from Bloomberg, Reuter and Expansión. I was very well-informed about global finance news. I was doing it just for fun. A few years went by and I decided to go to university thinking I would "learn more about finance". For sure I learnt a few new things, however I didn't like the whole university system.

After leaving university, I was kind of burnt out from finance. Hopefully I recovered my interest in this topic in a few months, and now I'm back.

![I'm back meme](/im-back.png)

After almost 30 years, I discovered what I want to do with my life.

Narrator: the previous statement is subject to changes.

I want to use my programming skills to make something that people use. If that thing is related to the finance sector, that would be awesome. TinyMacro is my first serious attempt to achieve this goal.

## So, what is TinyMacro?

Have you used the Bloomberg Terminal? My university had one on campus. That was probably the best decision they made. It's amazing how much data they have. In fact, this is what Bloomberg says about their systems:

"Our systems handle more than 300 billion ticks of data every day – in real time – ranging from live market data to news stories, financial tweets and instant messages. Our engineers build thousands of applications to not only process that information, but turn it into actionable insights for our clients."

Source: <https://www.bloomberg.com/company/values/tech-at-bloomberg/c-plus-plus/>

The Bloomberg Terminal is the best financial platform without any doubt, and they know it. That's because they charge you 25,000USD per year per user.

So what, Dennis? Are you trying to compite with Bloomberg? No. Shut up and listen.

### TinyMacro's main goal

Let's say you are a researcher that is writing an article about the correlation between inflation and sausage production in Germany. Or perhaps you are a retail investor interested in tracking the evolution of debt in the United States over the last 10 years. Maybe you're an Apple fanboy who wants to examine Apple's balance sheet. Wouldn't it be cool to have a single platfform to access this kind of data?

That's what I want to do with TinyMacro. I'm very bad at describing stuff but here is my attempt:

TinyMacro is a user-friendly web platform offering global macroeconomic data, emphasizing ease of use, data accuracy, and a standardized structure. It targets retail investors, academics, and hobbyists, making economic data more accessible.

Most global economic data is free and can be retrieved from government websites. Some examples are: SEC, Data.gov, World Bank, European Bank, National Statistical Institutes like the Federal Statistical Office in Germany, Eurostat, and so on.

My work is to automate data extraction, standardized structures for storage in a database, and continuously monitoring for any new changes. Then, build a web platform on top of that, and an API to offer the data to developers. It seems simple but it is not. It is a lot of work.

### What about stock prices?

What isn't available for free is real-time pricing for financial products. Once you start offering real-time pricing you'll find yourself competing with a lot of well-established big companies with multi-billion budgets. Physical location is also a competitive factor. If your data center is located close to the data providers (such as the exchanges), you can access data more rapidly, and I cannot have an office in Wall Street. Also, buying data directly from exchanges is expensive for me. That's why I cannot offer real-time pricing, yet.

And that is a brief introduction to what I've been working on over the past two weeks. Before wrapping up this post, I want to emphasize that it truly is a lot of work. Perhaps after a few months I realize that the idea is not viable for a solo-developer, or maybe not. We'll see.

Thanks for reading.

Dennis.
