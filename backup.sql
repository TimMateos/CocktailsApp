--
-- PostgreSQL database dump
--

\restrict XSY6Yq8NHFg6d0jziX8Hu7q1aXyObYh7iaDt97LLGO56JmHn9ylihjPxjLy4SNg

-- Dumped from database version 18.4
-- Dumped by pg_dump version 18.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: cocktails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cocktails (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    category character varying(100) NOT NULL,
    ingredients text NOT NULL,
    method text NOT NULL,
    serving text NOT NULL,
    is_iba boolean DEFAULT false,
    image_path text DEFAULT ''::text,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.cocktails OWNER TO postgres;

--
-- Name: cocktails_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cocktails_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cocktails_id_seq OWNER TO postgres;

--
-- Name: cocktails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cocktails_id_seq OWNED BY public.cocktails.id;


--
-- Name: inventory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inventory (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    in_stock boolean DEFAULT false
);


ALTER TABLE public.inventory OWNER TO postgres;

--
-- Name: inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inventory_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.inventory_id_seq OWNER TO postgres;

--
-- Name: inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inventory_id_seq OWNED BY public.inventory.id;


--
-- Name: cocktails id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cocktails ALTER COLUMN id SET DEFAULT nextval('public.cocktails_id_seq'::regclass);


--
-- Name: inventory id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory ALTER COLUMN id SET DEFAULT nextval('public.inventory_id_seq'::regclass);


--
-- Data for Name: cocktails; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cocktails (id, name, category, ingredients, method, serving, is_iba, image_path, created_at) FROM stdin;
12	Mojito	Refreshing	Белый ром - 45 мл\r\nСок лайма - 20 мл\r\nМята - 6-8 листьев\r\nСахар - 2 ч.л.\r\nСодовая или Спрайт	Смешать и подавить мяту с соком лайма и сахаром. Добавить ром и залить доверху содовой или спрайтом\r\nАккуратно перемешать	веточка мяты и долька лайма	t	/uploads/1781777550_Снимок экрана — 2026-06-18 в 13.08.18.png	2026-06-18 10:12:30.016147+00
13	South Side	Strong	Джин - 60 мл\r\nЛимонный сок - 30 мл\r\nСахарный сироп - 15 мл\r\nМята - 5-6 листьев\r\nЯичный белок (по вкусу)	Смешать все ингредиенты в шейкере со льдом и хорошо взболтать. Через двойное сито перелить в бокал\r\n\r\nЕсли делать в яичным белком, то взболтать нужно сильнее	веточка мяты	t	/uploads/1781777792_Копия iba-cocktail-new-era-south-side-6695d3bdcc9fe.jpg	2026-06-18 10:16:32.367209+00
6	French 75	Sour	Джин - 30 мл\r\nЛимонный сок - 15 мл\r\nСахарный сироп - 15 мл\r\nШампанское - 60 мл	Всё, кроме шампанского, влить в шейкер. Взболтать. Перелить в бокал\r\nДобавить шампанское и аккуратно размешать	N/A	t	/uploads/1781775589_Копия iba-cocktail-contemporary-classics-french-75-6695cdb175e06.jpg	2026-06-18 09:39:49.33612+00
7	Gin Fizz	Sour	Джин - 45 мл\r\nЛимонный сок - 30 мл\r\nСахарный сироп - 10 мл\r\nСодовая - 10 мл	Смешать все ингредиенты в шейкере со льдом. Взболтать\r\nПерелить в высокий бокал и добавить содовую	Долька лимона или цедра	t	/uploads/1781776058_Копия iba-cocktail-the-unforgettables-gin-fizz-6694910fc2eab.jpg	2026-06-18 09:47:38.252698+00
5	Cuba Libre	Refreshing	Белый ром - 50 мл\r\nКола - 120 мл\r\nСок лайма - 10 мл	Влить все ингредиенты в бокал хайболл, заполненный льдом	Долька лимона или лайма	t	/uploads/1781775412_Копия iba-cocktail-contemporary-classics-cuba-libre-6695cdb0a6f80.jpg	2026-06-18 09:36:52.224449+00
4	Daiquiri	Sour	Белый ром - 60 мл\r\nСок лайма - 20 мл\r\nСахар - 2 ложки	Смешать все ингредиенты в шейкере, размешивать ложкой до растворения сахара. Добавить лед и взболтать. \r\nПерелить в бокал	N/A	t	/uploads/1781775280_Копия iba-cocktail-the-unforgettables-daiquiri-6694910c5866e.jpg	2026-06-18 09:34:40.759174+00
8	John Collins	Sour	Джин - 45 мл\r\nЛимонный сок - 30 мл\r\nСахарный сироп - 15 мл\r\nСодовая - 60 мл\r\nЕсли использовать 'Old Tom' джин то получится Tom Collins	Влить все ингредиенты в бокал хайболл заполненный льдом и аккуратно перемешать	долька лимона и вишня	t	/uploads/1781776415_Снимок экрана — 2026-06-18 в 12.50.27.png	2026-06-18 09:53:35.300193+00
9	Long island	Refreshing	Текила - 15 мл\r\nДжин - 15 мл\r\nВодка - 15 мл\r\nБелый ром - 15 мл\r\nКуантро или Трипл Сек - 15 мл\r\nЛимонный сок - 25 мл\r\nСахарный сироп - 30 мл\r\nКола	Влить все ингредиенты в бокал хайболл заполненный льдом и залить колой доверху. Аккуратно перемешать	долька лимона	t	/uploads/1781776774_Копия iba-cocktail-contemporary-classics-long-island-iced-tea-6695cdc10c463.jpg	2026-06-18 09:59:34.008812+00
10	Margarita	Strong	Текила - 50 мл\r\nТрипл Сек - 20 мл\r\nСок лайма - 15 мл	Смешать все ингредиенты в шейкере со льдом и взболтать.\r\nПерелить в бокал	на краях бокала сделать кромку из соли	t	/uploads/1781776980_Копия iba-cocktail-contemporary-classics-margarita-6695cdd7505e0.jpg	2026-06-18 10:03:00.628687+00
11	Mint Julep	Strong	Бурбон - 60 мл\r\nМята - 3-4 веточки\r\nСахар - 1 ч.л.\r\nВода - 2 ч.л.	Аккуратно подавить мяту с сахаром и водой в бокале. Заполнить бокал колотым льдом, добавить Бурбон и аккуратно перемешивать до заледенения бокала	веточка мяты	t	/uploads/1781777226_Копия iba-cocktail-contemporary-classics-mint-julep-6695cdc4aa398.jpg	2026-06-18 10:07:06.124793+00
14	Tequila Sunrise	Sweet	Текила - 45 мл\r\nАпельсиновый сок - 90 мл\r\nСироп гренадин - 15 мл	В стакан хайболл заполненный льдом вылить текилу и апельсиновый сок. Сверху добавить сироп, не перемешивать	Апельсиновая долька либо цедра	t	/uploads/1781782711_Копия iba-cocktail-contemporary-classics-tequila-sunrise-6695cdd3da10a.jpg	2026-06-18 11:38:31.875074+00
15	White Lady	Sour	Джин - 40 мл\r\nТрипл Сек - 30 мл\r\nЛимонный сок - 20 мл	Смешать все ингредиенты в шейкере со льдом и хорошо взболтать. Перелить в охлажденный бокал	N/A	t	/uploads/1781782865_Копия iba-cocktail-the-unforgettables-white-lady-6694913318105.jpg	2026-06-18 11:41:05.388495+00
16	White Lady ©Krinkov	Sour	Белый ром - 40 мл\r\nТрипл Сек - 30 мл\r\nЛимонный сок - 20 мл	Смешать все ингредиенты в шейкере со льдом и хорошо взболтать. Перелить в охлажденный бокал	N/A	f	/uploads/1781783029_IMG_8200.jpeg	2026-06-18 11:43:49.68419+00
17	Martini Fiero	Sweet	Вермут фиеро - 120 мл\r\nГранатовый тоник - 200 мл\r\nАпельсин - одна-две дольки	Подавить апельсин и смешать все ингредиенты в большом бокале на ножке. Добавить лед и помешать	долька апельсина	f	/uploads/1781783410_foto-4448-11.jpg	2026-06-18 11:49:42.335216+00
18	Golden banana	Refreshing	Ром золотой - 50 мл\r\nБанановый сироп - 20 мл\r\nСок лайма или лимонный сок - 15 мл\r\nКола - 120 мл	В бокал хайболл заполненный льдом вылить ром, сироп и сок. Залить доверху колой и аккуратно перемешать	долька лимона, банана или веточка мяты	f	/uploads/1781783860_IMG_1377.jpeg	2026-06-18 11:57:41.002355+00
19	Daiquiri Banana	Sour	Белый ром - 80 мл\r\nБанановый сироп - 20 мл\r\nСахарный сироп - 15 мл\r\nСок лайма или лимонный сок - 60 vk	Смешать все ингредиенты в шейкере со льдом и хорошо взболтать, Процедить в бокал	банановая долька	f	/uploads/1781784078_3-retsept-koktejlya-dajkiri-630h368.jpg	2026-06-18 12:01:18.200902+00
20	Caribbean Blue (GPT ver)	Sweet	Белый ром - 40 мл\r\nБлю кюрасао — 20 мл\r\nСок лайма — 15 мл\r\nСок ананаса - 15 мл\r\nСахарный сироп - 10-15 мл (по вкусу)	Смешать все ингредиенты в шейкере и взболтать со льдом. Перелить в бокал хайболл	долька лимона	f	/uploads/1781784433_Копия cocktail-caribbean-blue.jpg	2026-06-18 12:07:13.443124+00
21	Cosmonaut	Sweet	Водка - 60 мл\r\nЛимонный сок - 20 мл\r\nБлю Кюрасао - 15 мл\r\nАнанасовый сок - 120 мл	В заполненный льдом бокал влить водку и сок, долить доверху ананасовый сок. Сверху добавить Блю Кюрасао и аккуратно размешать	лимонная долька	f	/uploads/1781785549_foto-3571-6.jpg	2026-06-18 12:16:58.998135+00
\.


--
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inventory (id, name, in_stock) FROM stdin;
24	Бурбон	f
25	Виски	f
1	Джин	f
2	Белый ром	f
3	Лимонный сок	f
4	Сок лайма	f
5	Водка	f
6	Золотой ром	f
7	Темный ром	f
8	Текила	f
9	Апельсиновый сок	f
10	Ананасовый сок	f
11	Банановый сироп	f
12	Банановый ликер	f
13	Сухой вермут	f
14	Сладкий вермут	f
15	Вермут Фиеро	f
16	Шампанское	f
17	Сахарный сироп	f
19	Апероль	f
20	Мята	f
21	Мятный сироп	f
22	Сироп гренадин	f
23	Блю Кюрасао	f
26	Содовая	f
27	Кола	f
28	Спрайт	f
29	Яблочный сок	f
30	Куантро	f
31	Трипл Сек	f
32	Амаретто	f
33	Пряный ром	f
34	Вода	f
35	Гранатовый тоник	f
36	Апельсин	f
37	Яичный белок	f
\.


--
-- Name: cocktails_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cocktails_id_seq', 21, true);


--
-- Name: inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.inventory_id_seq', 37, true);


--
-- Name: cocktails cocktails_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cocktails
    ADD CONSTRAINT cocktails_pkey PRIMARY KEY (id);


--
-- Name: inventory inventory_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_name_key UNIQUE (name);


--
-- Name: inventory inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

\unrestrict XSY6Yq8NHFg6d0jziX8Hu7q1aXyObYh7iaDt97LLGO56JmHn9ylihjPxjLy4SNg

