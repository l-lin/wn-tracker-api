CREATE EXTENSION "uuid-ossp";
CREATE TABLE default_novels (
  id        VARCHAR DEFAULT uuid_generate_v4(),
  title     VARCHAR NOT NULL,
  url       VARCHAR NOT NULL,
  feed_url  VARCHAR,
  image_url VARCHAR,
  summary   TEXT,
  favorite  BOOL
);

INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Overlord',
  'http://skythewood.blogspot.sg/p/knights-and-magic-author-amazake-no.html',
  'https://www.blogger.com/feeds/4225665848364118863/posts/default',
  'http://4.bp.blogspot.com/-c0DBPG-s_28/VEb8GzDZ2yI/AAAAAAAAApQ/tcl9rCbAmoQ/s1600/1.jpg',
  'After announcing it will be discontinuing all service, the internet game "Yggdrasil" shut downs -- That was the plan. But for some reasons, the player character did not log out some time after the server was closed. NPC starts to become sentient. A normal youth who loves gaming in the real world seemed to have been transported into an alternate world along with his guild, becoming the strongest mage with the appearance of a skeleton, Momonga. He leads his guild "Ainz Ooal Gown" towards an unprecedented legendary fantasy adventure!',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Altina',
  'http://skythewood.blogspot.sg/p/altina-sword-princess.html',
  'https://www.blogger.com/feeds/4225665848364118863/posts/default',
  'http://1.bp.blogspot.com/-AyC6RgreYRM/VDqYzsqFRvI/AAAAAAAAAj8/V6Ljf9LlxOY/s1600/a.JPG',
  'Unskilled in both swords and bows, Regis is a soldier who ranks at the bottom in military academy who is obsessed with books. After being banished to the borders, he met a girl who changed his destiny. With red hair and crimson eyes, Princess Altina who holds the sword of kings. Although she is a daughter of a concubine, she was appointed as the commander of the border regiment at the tender age of 14. But she did not lament her circumstances and bears higher aspiration. "I believe in you." Regis received the request of the girl to be her strategist and forge through many hardships. The sword princess and the biblophile youth embark on their fantasy war story --',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Knight''s and magic',
  'http://skythewood.blogspot.sg/p/blog-page.html',
  'https://www.blogger.com/feeds/4225665848364118863/posts/default',
  'http://4.bp.blogspot.com/-DhrMa-Qsi9s/VMjI3c-ZaYI/AAAAAAAABDg/mlN8BPNZ1EI/s1600/414px-Knights_and_Magic_v1_Cover.jpg',
  'A young man from Japan passed away after a traffic accident. His soul was reincarnated in an alternate world into the body of a pretty young boy Ernesti Echevarria with his memories intact. Influenced by his hobby from his previous life, Eru is a ''robot nerd'' in this life too. He meets the giant humanoid battle robots in this world -- Silhouette Knights. The elated Eru started a series of plans in order to pilot the robots. He drags his childhood friend in this world along as he messes around in this world to satisfy his desire for robots.',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Arifureta',
  'http://japtem.com/projects/arifureta-toc/',
  'http://japtem.com/feed/',
  'http://puu.sh/auzGa/fded6287f0.png',
  'Among the class transported to another world, Nagumo Hajime is an ordinary male student who didn''t have ambition nor aspiration in life, and thus called “Incompetent” by his classmates. The class was summoned to become heroes and save a country from destruction. Students of the class were blessed with cheat specifications and cool job class, however, it was not the case with Hajime, with his profession as a “Synergist”, and his very mediocre stats. “Synergist”, to put it in another word was just artisan class. Being the weakest, he then falls to the depth of the abyss when he and his classmates were exploring a dungeon. What did he find in the depth of the abyss, and can he survive?',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Ark',
  'https://arkmachinetranslations.wordpress.com/',
  'https://arkmachinetranslations.wordpress.com/feed/',
  'http://japtem.com/wp-content/uploads/2013/06/Arkrelease.png',
  'Kim Hyun Woo lived the life of the wealthy thanks to his parents. But one day, he received a phone call informing him of a traffic accident which involved his parents. His father had died and his mother was hospitalized in critical condition. The normal life he once knew, collapsed… They sold their house, cancel various insurance plans, and moved to a one room apartment. And after a few years, Hyun Woo spends four to six hours tending to his mother and worked to pay for her medical bills. One day, one of his Instructors recommended him for position in a company called Global Exos, who made an announcement of making history with the newest technological invention. This story follows the main protagonist on his journey to adapt to a new development of a virtual reality game.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'World teacher',
  'https://defiring.wordpress.com/',
  'https://defiring.wordpress.com/feed/',
  '',
  'A man who was once called the world''s strongest agent ended up becoming a teacher after his retirement, to train the new generation of agents. After many years of training his disciples, he was killed at the age of 60 by the ploy of a secret organization and was reincarnated in another world with all his past memories. Though he was surprised by the magic and the strange species of that world, he adapted very fast to his condition as a newborn and took advantage of it. He acquired special magic and gained a massive amount of strength thanks to his tight discipline, in order to reach his goal: Resume his career as a teacher which he left halfway through in his previous life. This is the story of a man, who, based on the memories and the experiences of his previous life, became a teacher who travels through the world with his students.',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Legendary moonlight sculptor',
  'http://jawztranslations.blogspot.fr/p/legendary-moonlight-sculptor.html',
  'http://jawztranslations.blogspot.com/feeds/posts/default',
  'http://japtem.com/wp-content/uploads/2013/06/lmsrelease.png',
  'Lee Hyun, a debt-ridden teenage orphan, puts his account of the top-level avatar in the Continent of Magic up for auction. Once the news is spread nationwide, a fierce competition to win his avatar breaks out, making him a multimillionaire overnight. However, when men in black show up to collect the debts Hyun''s parents owed them; he ends up where he started — almost penniless. Determined to exact revenge on the loan sharks, he sets out a journey with Weed as his alias, the lowliest of the lowly, in the online RPG game named Royal Road to prove history can repeat itself. WHY NOT CLAIM THE TOP ONCE AGAIN? Weed''s sugary words and tough action work magic one after another — only until he finds himself becoming a Sculptor, the most contemptible profession in terms of combat ability and career prospect…',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Zectas',
  'http://jawztranslations.blogspot.fr/p/zectas.html',
  'http://jawztranslations.blogspot.com/feeds/posts/default',
  '',
  'At 18 years old, Nash Smoak maintains three jobs. He works as an assistant to the cook in a small diner, he is shadily hired as a construction worker, and works as an attendant at a game arcade. Nash needs to take on those jobs in order to support his two brothers and sickly Grandma. Ever since the tragic accident of his parent''s death. He took it upon himself to support his family. After being worked to the bone, Nash found that sacred place, his paradise of relief, the Virtual Reality World of Zectas. However, his only sanctuary was destroyed when a group of high level players used him as bait in one of their quests. Nash was captured and tortured by the Legendary Monster! Little did they know that this incident would cause the birth of the Legendary User Smoke!',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Drezo regalia',
  'http://jawztranslations.blogspot.fr/p/drezo-regalia.html',
  'http://jawztranslations.blogspot.com/feeds/posts/default',
  'https://lh5.googleusercontent.com/vgsmfqMS9EAs3Jvn6OjRqN5NgSCbLNSdyjqe1wonscUT6KSVtRlKOlT9ATHOgHgfXuDu30zVoX2LmrT0vQXQea7BiwejqyPQzbltETnQrQ-dwf5DoRHaXxvlOvqDRVMsow',
  'Yami Hikari company had created a new game gear called the Embyro that are split into two different gears: the first is the deep dive gear called ‘Sense'' and the second is high tech glasses called ‘Alive.'' One is made to transport the players mind to a virtual world, while the other is the next generation new technology that incorporates the reality with the game. A game where it test the limits of technology and human boundaries. What is real? What is fake? Who is right or wrong? What is the Seed? Is life just a game? Three different people''s destinies are intertwined: One is on a mission, One wants change in life, and One wants to be left alone. Their story converges in a Virtual Reality MMO game called Growth. This is a story of Zero, Agnis, and BlackStar.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Sendai Yuusha wa Inkyou Shitai',
  'https://manga0205.wordpress.com',
  'https://manga0205.wordpress.com/feed/',
  '',
  'The Summoned Hero and the Preceding Hero',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Slave harem',
  'https://wartdf.wordpress.com/',
  'https://wartdf.wordpress.com/feed/',
  '',
  'Slave harem in the labyrinth of the other world',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Seirei Gensouki',
  'https://zmunjali.wordpress.com/category/seirei-gensouki-konna-sekai-de-deaeta-kimi-ni/',
  'https://zmunjali.wordpress.com/feed/',
  '',
  'Amakawa Haruto, a young man who died before reuniting with his childhood friend who disappeared five years ago. Rio, a boy living in the slums who wants revenge for his mother who was killed when he was five years old. Earth and another world. Two people who have completely different backgrounds and values, but for some reason, the memories and personality of Haruto who should''ve died resurrected in Rio''s body. As the two are confused over their memories and personalities fusing together, Rio (Haruto) decide to live in this new world. For some reason, along with Haruto''s memories, Rio awakened an unknown "special power," and it seems that if used well, he can live a better life than he does now. But before that, Rio encountered the princess kidnapping incident of the Bertram Kingdom he lives in. Saving the princesses, Rio was accepted into the state managed Royal Academy. Being an orphan in a school for nobles, it was an extremely destestable place to him.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Coiling Dragon',
  'http://www.wuxiaworld.com/cdindex-html/',
  '',
  'https://www.mangaupdates.com/image/i206448.jpg',
  'Linley is a young noble of a declining clan. He has large aspirations and wants to save his clan. Linley''s journey begins with an accident when he discovers a ring. He took a liking to this ring with a dragon coiling around its entirety. Upon being injured during a battle between two powerful fighters he discovers that his ring is not what he thought it was and possesses powers beyond his imagination.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Risou no Himo Seikatsu',
  'http://unlimitednovelfailures.mangamatters.com/risou-no-himo-seikatsu/',
  'http://unlimitednovelfailures.mangamatters.com/feed/',
  'http://unlimitednovelfailures.mangamatters.com/wp-content/uploads/2014/07/Cover.jpg',
  'Summoned by a beautiful woman to a different world, Yamai Zenjirou is asked to marry her and make a child with her. Will he throw his life on earth away for a sponger life with a beautiful woman?',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Konjiki no Wordmaster',
  'https://yoraikun.wordpress.com/knw-chapters/',
  'https://yoraikun.wordpress.com/feed/',
  'https://i0.wp.com/cdn.myanimelist.net/images/manga/3/142615l.jpg',
  'The Unique Cheat of the Man Dragged in by the Four Heroes',
  TRUE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Tate No Yuusha No Nariagari',
  'https://yoraikun.wordpress.com/translated-chapters/',
  'https://yoraikun.wordpress.com/feed/',
  'https://i0.wp.com/static.zerochan.net/Tate.No.Yuusha.No.Nariagari.full.1768808.jpg',
  'The Rise of the Shield Hero',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Suterurareta Yuusha no Eiyuutan',
  'https://pirateyoshi.wordpress.com/translated-chapters/',
  'https://pirateyoshi.wordpress.com/feed/',
  '',
  'Katsuragi is an average, overweight otaku who''s bullied in class. One day when he''s bullied their whole class is summoned by a beautiful goddes to go to another world to become heroes. They were promised any wish, and our "hero", because goddess was beautiful, he wanted her, so he made the whole class accept her offer. They were welcomed by a king of some country. They were received with a hero''s welcome, but when they got to see their statuses, they immediately treated Katsuragi like s**t, since he was very, very average. When they entered the dungeon to level up, and after falling for a trap (Too much Arifureta...), our MC was sacrificed (ie. "Hero" threw him because his magic, Katsuragi''s, was depleted, and they needed a decoy, and who''s better than an useless, magic depleted, Katsuragi. Katsuragi was eaten by monsters, but after some time he regained his consciousness... And after checking his status window, he noticed that his stats skyrocketed (easily 4-6 times higher than the "hero"), and he noticed he had a lot of Special Abilities, which he previously lacked (he was the only one in his class who lacked any special abilities), like: Hearth of Steel, 1/3 chances of preventing poisoning, hypnosis, paralysis and mind control magic (???), Magic is not consumed when he reaches 100 in magic (he currently has 2520), King of Death which means he could revive people who recently died, and some stuff about him reviving, and revenge abilities (GT is VERY unclear about this). He also noticed that he died 5 times! After waking up, he goes to some door and notices some kind of silver wolf that the "hero" couldn''t defeat, and he defeated him in one punch through the brain... Then he noticed his classmate who was dead. He revived her. She also has big oppai',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Suterurareta Yuusha no Eiyuutan',
  'https://oniichanyamete.wordpress.com/index/maou/',
  'https://oniichanyamete.wordpress.com/feed/',
  '',
  'Demon King Luruslia Nord; defeated by the hero, and someone who should have died. He had a number of things left undone, and a number of problems left unresolved, but following the slogan that ‘evil will be destroyed'', he should have disappeared. However, for some reason, he had miraculously gained a second life as a human.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'RE:Monster',
  'https://docs.google.com/document/d/1t4_7X1QuhiH9m3M8sHUlblKsHDAGpEOwymLPTyCfHH0/preview',
  '',
  'http://img.batoto.net/forums/uploads/8821ef3fcf8b23aaed10a45e3aa7b34a.jpg',
  'Re:Monster is about a main character who gets stabbed and killed, and is then born again in a fantasy world… as a Goblin. He inherits his ESP Ability “Absorption” by eating various things to learn new skills and abilities.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Tsuyokute New Saga',
  'https://docs.google.com/document/d/17taISyRebqOdRhIuclyBT94kKbWrJgN3Fc1WvyOWey0/preview',
  '',
  'https://tensaitranslations.files.wordpress.com/2014/08/tns.jpg?w=211&h=314',
  'Kail and his band of suicide squad head straight into Maou''s Castle to put an end to his schemes. With him losing his right arm and leg, he vanquishes Maou at the cost of his team and his loved one… Slowly dying as he bleeded out, he saw a dark crystal in the area, deciding to grab it. A blinding light surrounded him… and somehow he finds himself on a bed in a familiar room.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'EC',
  'http://ecwebnovel.blogspot.ca/',
  'http://ecwebnovel.blogspot.com/feeds/posts/default',
  '',
  'In the distant future, after humanity''s revival following an apocalyptic event, an inquisitive woman, a young man abandoned by his country, a youth seeking redemption and an isolated young prodigy meet by chance in a game promoted by a certain company.  Tricked into running an Academy, the tale follows their lives and adventures in a world not their own.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Master of Monster',
  'https://docs.google.com/document/d/1DL6dpctSl_DME-4cnnZ8mkfCIb9nIMlu8P8JhDU-it8/preview',
  '',
  'http://vignette3.wikia.nocookie.net/monster-tamer/images/6/6f/MT01.jpg/revision/latest/scale-to-width/300?cb=20140818114424',
  'An entire class has been tossed into a fantasy world without knowing what happened. Suddenly monsters attack and kill a lot of students, but then other students started to fight back with cheat abilities. A few days later, everyone has been divided into the Stay Home Group and Exploration Group. The exploration group wanted to see if there''s anything else to find and take care off in the forest that they are in. They are armed to brim with cheats and makeshift weapons. The Stay Home Group doesn''t have cheat abilities…. after the Exploration Group leaves, the stay home group gets split as a riot ensues. Student killing student, girls getting raped by their fellow classmates… and our MC starts to run towards the forest. After running, evading lots of monsters… he wounds up in a cave, exhausted and almost ready to die… when a Slime approaches him. It begins to eat his hand and all he can do was plead "...please… someone out there… help me…."',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Gun-Ota ga Mahou Sekai ni Tensei',
  'https://docs.google.com/document/d/1Lzfs2R1eYPfBDoMEXP4NKsorA7TTQD1Yl9JXZsShAaM/pub',
  '',
  '',
  'Hotta Yotta is killed while returning home from work on a cold evening. He wakes up being carried by a bunny-eared woman who is a body of a baby! Featuring Reincarnation, Guns and Harem, its a Gun vs Magic Story that gives Hotta Yotta a new life… and an edge!',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Souen no Historia',
  'https://docs.google.com/document/d/1PfyUTq0EfhjdRABQFLbC2C9ByHTzm1CcsnDxq2h3t8c/preview',
  '',
  '',
  'Dying after saving a girl, protagonist wakes up sucking boobs. He is now a baby! But he remembers everything from his past… Now his newself is a butler for an ojou-sama. This is a story of a reincarnation turned into a Super Butler that would make Hayate (from Hayate no Gotoku) piss his pants and cry.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Slime Tensei Monogatari',
  'https://docs.google.com/document/d/10KAGFVd2ESbuuALG--fxFbyEALfOfS2t0YjL2GlIkaY/preview',
  '',
  '',
  'An old office worker dies… and is reincarnated as a Slime! Watch as he survives a hard and dangerous world with his wits!',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Yuusha Party no Kawaii ko ga ita no de, Kokuhaku Shite Mita',
  'https://docs.google.com/document/d/1rhoCQcOUW1eccV4Fsh6K76u0Vup7DMQhBfDcknEbFHI/preview',
  '',
  '',
  'A boy, Youki, is reincarnated as a demon mini-boss in a fantasy world, but when the Hero Party comes to kill the Demon Lord, he falls in love with the cute cleric Cecilia. This love at first sight causes a problem as he is no longer a human, but part of the demon enemy she was sent to defeat. What on Earth will happen when he tries confessing to her?',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Shinka no Mi',
  'https://docs.google.com/document/d/1P8975xchjaZNw7gEiQkrV1w01akcJhUjlcd4fo4uWNU/edit',
  '',
  '',
  'Hiiragi Seiichi is an ugly, revolting, dirty, smelly fatass; these are the insults hurled at him one after another about his appearance. Such was Seiichi''s daily school life of bullying, then for some reason, one day when school was out, a voice claiming to be a God said over the PA system to prepare to be transported to another world. What''s more, not Seiichi alone, but the entire school. A fantasy world where game-like elements such as as levels, stats, and skills exist.  However, the God still had preparations to complete for the transfer, and would send them over as soon as the hero summoning ritual was ready. The classes all formed groups to wait for the transfer, but Seiichi alone was discluded and as such was summoned to a different area. After being transported the first thing Seiichi ate was the "Fruit of Evolution". This would come to greatly change his life ---- This story is about how Seiichi went from being severely bullied by his classmates, even not being recognized for his accomplishments, and despite all that staying positive and surviving in this new world. As a result, he somehow becomes one of the champions.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Jashin Tensei',
  'https://docs.google.com/document/d/10_3RfB9ZSNhNrlsF1YB53NJqjQ1Nu3F0b6d01bfPHxU/preview',
  '',
  'http://ecx.images-amazon.com/images/I/51RT3mQgXkL.jpg',
  'After going to the great beyond, Hirano Bonta was reincarnated as an Evil God. Descending upon the Demon Lands, he came across his follower, the 108th Demon Lord, the <Dethroned Crown Prince> Drake. A fortuitous omen...thought Drake, who had suffered a great defeat in battle, leaving him with only 200 underlings. On top of that, pursuers were hot on his tail; he was between a rock and a hard place. But the genius strategist Drake was not ready to give up just yet. Despite the crisis, he still held the ambition of unifying his nation once again. The history-geek Hirano was super excited and was determined to support Drake as an Evil God, but as a newbie Evil God there was little he could do...just how will they get through this together?!',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Eh? Heibon Desu yo??',
  'https://docs.google.com/document/d/1xInAD8v06AIX_urMZRRXHBocDsqBEePMoU1EOTfGRZQ/pub',
  '',
  '',
  'Yukari was a high school student when she died in a traffic accident, but when she woke up, she had been reincarnated as the daughter of a count in another world! But strangely, what was waiting for her was a life of poverty, so she decided to make use of the knowledge from her previous life.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Sayonara Ryuusei Konnichiwa Jinsei',
  'https://binhjamin.wordpress.com/sayonara-ryuusei-konnichiwa-jinsei/',
  '',
  '',
  'The oldest and strongest Dragon grew tired of living choose to die when the Heroes come for his life. When the Dragon''s soul is waiting to drift toward the Sea of Soul, that is when its noticed its has been reborn into a Human baby. The Dragon then decide to live as a life as a Human to its fullest, and regained his will to live. The Dragon was born into the child of a farmer lives his live in the frontier possessing enormous amount of power due to his soul being a Dragon''s. He then encounter the Demon of the Lamia race, fairies, Black Rose, and later enter the Magic Academy. The man whose soul is of a Dragon lives his life in joy going to Magic Academy, spending time with an old friend, The Earth Goddess, meets beautiful girls, strong classmates, the Dragon King and the Queen of Vampires.',
  FALSE
);
INSERT INTO default_novels (title, url, feed_url, image_url, summary, favorite)
VALUES (
  'Bu ni Mi wo Sasagete Hyaku to Yonen',
  'https://binhjamin.wordpress.com/bu-ni-mi-wo-sasagete-hyaku-to-yonen/',
  'https://binhjamin.wordpress.com/feed/',
  'https://binhjamin.files.wordpress.com/2014/12/img_0001.png?w=213&h=300',
  'Slava, who devoted his entire life practicing martial arts, took a disciple from a foreign country. The undefeated martial artist, never losing a fight up until now, has been overcome by old age is about to take his last breath over a century after his birth. While lying in bed regretting his past, Slava noticed his disciple caring for him. The man and small woman observe each other. She says “Yoshi, Papa~” which causes Slava to be confused with a single word. The man who regrets his lifelong devotion to martial arts, reincarnated into the race which is known for vast vitality, the Elven race. Using the knowledge from his previous life, he gleefully trained his body and… ?',
  FALSE
);

CREATE TABLE novels (
  id           VARCHAR   DEFAULT uuid_generate_v4(),
  token        VARCHAR NOT NULL,
  title        VARCHAR NOT NULL,
  url          VARCHAR NOT NULL,
  feed_url     VARCHAR,
  image_url    VARCHAR,
  summary      TEXT,
  favorite     BOOL,
  last_updated TIMESTAMP DEFAULT now()
);
