--
-- PostgreSQL database dump
--

-- Dumped from database version 15.7 (Ubuntu 15.7-1.pgdg20.04+1)
-- Dumped by pg_dump version 15.8 (Debian 15.8-1.pgdg120+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: heroku_ext; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA heroku_ext;


--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

-- *not* creating schema, since initdb creates it


--
-- Name: pg_stat_statements; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA heroku_ext;


--
-- Name: EXTENSION pg_stat_statements; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pg_stat_statements IS 'track planning and execution statistics of all SQL statements executed';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: categories; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.categories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    type text,
    parent_id bigint
);


--
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- Name: certifications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.certifications (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    logo text,
    industry text,
    certifier text,
    certifies_companies boolean,
    certifies_products boolean,
    certifies_processes boolean,
    certifier_contact_id text,
    audited boolean,
    auditor text,
    region text,
    qualifiers text,
    sources text,
    website text,
    description text,
    certifies_company boolean,
    certifies_product boolean,
    certifies_process boolean
);


--
-- Name: certifications_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.certifications_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: certifications_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.certifications_id_seq OWNED BY public.certifications.id;


--
-- Name: companies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.companies (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    "name json:name" text,
    url text,
    description text,
    user_id bigint,
    is_verified boolean,
    city text,
    state text,
    country text,
    image_id bigint,
    image text
);


--
-- Name: companies_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.companies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: companies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.companies_id_seq OWNED BY public.companies.id;


--
-- Name: favourites; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.favourites (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_refer bigint,
    product_id bigint
);


--
-- Name: favourites_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.favourites_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: favourites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.favourites_id_seq OWNED BY public.favourites.id;


--
-- Name: images; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.images (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    image_location text
);


--
-- Name: images_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.images_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: images_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.images_id_seq OWNED BY public.images.id;


--
-- Name: login_credentials; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.login_credentials (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email text,
    password_hash text,
    salt text
);


--
-- Name: login_credentials_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.login_credentials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: login_credentials_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.login_credentials_id_seq OWNED BY public.login_credentials.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    url text,
    description text,
    title text
);


--
-- Name: products_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: products_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    login_credentials_id bigint
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- Name: certifications id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.certifications ALTER COLUMN id SET DEFAULT nextval('public.certifications_id_seq'::regclass);


--
-- Name: companies id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.companies ALTER COLUMN id SET DEFAULT nextval('public.companies_id_seq'::regclass);


--
-- Name: favourites id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.favourites ALTER COLUMN id SET DEFAULT nextval('public.favourites_id_seq'::regclass);


--
-- Name: images id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images ALTER COLUMN id SET DEFAULT nextval('public.images_id_seq'::regclass);


--
-- Name: login_credentials id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.login_credentials ALTER COLUMN id SET DEFAULT nextval('public.login_credentials_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.categories (id, created_at, updated_at, deleted_at, name, type, parent_id) FROM stdin;
\.


--
-- Data for Name: certifications; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.certifications (id, created_at, updated_at, deleted_at, name, logo, industry, certifier, certifies_companies, certifies_products, certifies_processes, certifier_contact_id, audited, auditor, region, qualifiers, sources, website, description, certifies_company, certifies_product, certifies_process) FROM stdin;
15	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	BIFMA LEVEL	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285565/BIFMALEVEL.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.bifma.org/page/level	Certifies furniture products based on environmental impact, health and wellness, and social responsibility criteria.	f	t	\N
16	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Blue Angel	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285566/BlueAngel.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.blauer-engel.de/en	Guarantees that a product meets high environmental standards, including protecting consumers' health.	f	t	\N
17	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Bluesign	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285552/Bluesign.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.bluesign.com/en	Products are responsibly manufactured using safer chemicals and fewer resources, including less energy, in production.	f	t	\N
18	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Carbon Neutral by Carbon Trust	https://res.cloudinary.com/dwmpianpg/image/upload/v1707284639/CarbonNeutralByCarbonTrust.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.carbontrust.com/	Products reduce their carbon footprint year after year and any outstanding emissions are offset.	f	t	\N
20	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Carbonfree Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285546/CarbonfreeCertified.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://carbonfund.org/	Determines the carbon footprint of the product and associated carbon emissions are offset with reduction projects.	f	t	\N
21	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	CarbonNeutral product by Climate Impact Partners	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285547/CarbonNeutralproductbyClimateImpactPartners.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.carbonneutral.com/	Measures all the emissions created in the product’s manufacture, makes internal reductions and offsets the remainder.	f	t	\N
22	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Climate neutral by ClimatePartner	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285544/ClimateNeutralbyClimatePartner.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.climatepartner.com/en	Verifies the product’s carbon footprint was calculated, is continuously reduced and remaining emissions were offset.	f	t	\N
23	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Compact by Design (Amazon developed)	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285545/CompactbyDesign(Amazondeveloped).jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.amazon.com/b?node=21221609011	Products remove excess air and water, which reduces the carbon footprint of shipping and packaging.	f	t	\N
24	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Cradle to Cradle Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285548/CradletoCradleCertified.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.c2ccertified.org/	Products are made with safer materials and responsible processes to positively impact people and our planet. Includes Bronze, Silver, Gold, and Platinum levels.	f	t	\N
25	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	ECOLOGO	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285566/ECOLOGO.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.ul.com/resources/ecologo-certification-program	Certified products meet standards that can reduce the environmental impact of one or more stages of the product lifecycle.	f	t	\N
19	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	Carbon Trust	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285546/CarbonTrustLogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.carbontrust.com/	Reducing CO2 products reduce their carbon footprint year after year. Certified by the Carbon Trust.	f	t	\N
26	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	ENERGY STAR Most Efficient	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285557/ENERGYSTARMostEfficient.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.energystar.gov/products/most_efficient	Recognizes the best of ENERGY STAR certified products with highest efficiency and maximum carbon reductions.	f	t	\N
27	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	EPA Safer Choice	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285551/EPASaferChoice.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.epa.gov/saferchoice	U.S. EPA Safer Choice certified products contain safer ingredients for human health and the environment.	f	t	\N
28	2024-02-10 18:58:20.430794+00	2024-02-10 18:58:20.430794+00	\N	EPEAT	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285557/EPEAT.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://globalelectronicscouncil.org/ecolabels/	Products are assessed against criteria including energy use and have a reduced sustainability impact across their lifecycle. Includes Silver and Gold ratings.	f	t	\N
29	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	EU Organic	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285553/EUOrganic.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://ec.europa.eu/info/food-farming-fisheries/farming/organic-farming/organic-logo_en	Products have zero or minimal chemical pesticides or fertilizers, support animal welfare and non-genetically modified standards.	f	t	\N
30	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	EWG Verified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285553/EWGVerified.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.ewg.org/ewgverified/	Products are reviewed to ensure they are free from EWG’s known chemicals of concern and adhere to strict health standards.	f	t	\N
31	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Fair for Life	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285559/FairforLife.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.fairforlife.org/	Certifies for fair trade and more socially and environmentally responsible supply chains.	t	t	\N
32	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Fair Rubber	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285560/FairRubber.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://fairrubber.org/	Products contain natural rubber that is certified as fairly traded according to social and environmental criteria.	f	t	\N
33	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Fair Trade Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285558/FairTradeCertified.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.fairtradecertified.org/	Products are made according to standards that promote safety, sustainability, and empowerment in the workplace.	t	t	\N
34	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Fairtrade International	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285560/FairtradeInternational.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.fairtradeamerica.org/why-fairtrade/	Products are produced in line with ethical and environmental standards, including supporting farmers to tackle climate change challenges.	f	t	\N
35	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Global Organic Textile Standard	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285568/GlobalOrganicTextileStandard.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.global-standard.org/	Certifies each step of the organic textile supply chain against strict ecological and social standards.	f	t	\N
36	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Global Recycled Standard	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285564/GlobalRecycledStandard.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/recycled-claim-standard-global-recycled-standard/	Products are made with at least 50% recycled content and meet social, environmental, and chemical requirements.	f	t	\N
37	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	Green Seal Logo	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285569/GreenSealLogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://greenseal.org/	Green Seal certified products have reduced climate and environmental impacts at one or more stages of their lifecycle.	f	t	\N
38	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	GreenCircle Certified: Certified Energy Savings	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285558/GreenCircleCertified:CertifiedEnergySavings.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.greencirclecertified.com/certified-energy-savings-amazon-cpf	Verifies product energy savings of over 10% in the use-phase compared to similar products.	f	t	\N
39	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	GreenCircle Certified: Certified Environmental Facts Label	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285546/GreenCircleCertified:CertifiedEnvironmentalFactsLabel.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.greencirclecertified.com/certified-environmental-facts-amazon-cpf	Verifies product carbon footprint reduction in the raw material or manufacturing stage.	f	t	\N
40	2024-02-10 19:06:16.62955+00	2024-02-10 19:06:16.62955+00	\N	GreenCircle LCA Optimized	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285568/GreenCircleLCAOptimized.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.greencirclecertified.com/lca-optimized-amazon-cpf	GreenCircle Certified: Life Cycle Assessment Optimized verifies reductions in product life cycle impacts through product or manufacturing improvements.	f	t	\N
41	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	ISCC PLUS	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285548/ISCCPLUS.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.iscc-system.org/certification/iscc-certification-schemes/iscc-plus/	ISCC PLUS certified products are created using raw materials like recycled plastics and bio-based materials, which can reduce fossil fuel consumption.	f	t	\N
42	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Leather Working Group	https://res.cloudinary.com/dwmpianpg/image/upload/v1707284585/leather-working-group.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.leatherworkinggroup.com/	Leather Working Group certification promotes more-responsible leather production and is committed to driving positive environmental change.	t	t	\N
43	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Made Safe	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285549/MadeSafe.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.madesafe.org/	MADE SAFE certifies products made without known toxic chemicals, by prohibiting substances that harm human health or the environment.	f	t	\N
44	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Natrue	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285569/Natrue.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.natrue.org/	Natrue certified cosmetics are made from naturally sourced ingredients.	f	t	\N
45	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Nordic Swan Ecolabel	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285570/NordicSwanEcolabel.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.nordic-ecolabel.org/	Nordic Swan Ecolabel certified products meet requirements on chemical usage, resource consumption, greenhouse gas emissions, and biodiversity.	f	t	\N
46	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	OEKO-TEX MADE IN GREEN	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285549/OEKOTEXMADEINGREEN.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.oeko-tex.com/en/our-standards/made-in-green-by-oeko-tex	OEKO-TEX MADE IN GREEN products are tested for harmful substances and made in safer workplaces with reduced environmental impacts.	f	t	\N
47	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	OEKO-TEX STANDARD 100	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285550/OEKOTEXSTANDARD100.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.oeko-tex.com/en/our-standards/standard-100-by-oeko-tex	OEKO-TEX STANDARD 100 requires textiles-based products to be tested against a list of 1,000+ chemicals, to avoid those which may be harmful to human health.	f	t	\N
48	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Organic Blended Content Standard	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285554/OrganicBlendedContentStandardlogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/organic-content-standard/	Organic Content Standard Blended products included in the program are made with at least 50% organically grown materials.	f	t	\N
49	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Organic Content Standard 100	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285554/OrganicContentStandard100.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/organic-content-standard/	Organic Content Standard 100 products are made with at least 95% organically grown materials.	f	t	\N
50	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Plant-Based Fiber Blended	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285555/Plant-BasedFiberBlended.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.amazon.com/b/?node=96704931011	Plant-Based Fiber Blended (Amazon developed) products are made with at least 50% plant-based content and produced in a way that restricts the use of harmful chemicals.	f	t	\N
51	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Pre-owned Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285564/Pre-ownedCertifiedlogo.png	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.amazon.com/s/browse/?node=23911980011	Pre-owned Certified: Electronics (Amazon-developed Certification) products are inspected, cleaned, and (if applicable) repaired to excellent functional standards. Buying Pre-owned extends a product's life, reducing e-waste and raw material extraction.	f	t	\N
52	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Preowned Certified Fashion	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285565/PreownedCertifiedFashionLogo.png	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.amazon.com/b?node=109618761011	Pre-owned Certified: Fashion (Amazon-developed Certification) products have been previously owned. Buying pre-owned instead of new can reduce greenhouse gas emissions, the use of fossil fuels and water, and waste to landfill.	f	t	\N
53	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Rainforest Alliance	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285560/RainforestAlliancelogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.rainforest-alliance.org/faqs/what-does-rainforest-alliance-certified-mean	Rainforest Alliance’s seal stands for more-sustainable farming methods that help improve farmer livelihoods & mitigate climate change.	t	t	\N
54	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Recycled Claim Standard	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285562/RecycledClaimStandard.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/recycled-claim-standard-global-recycled-standard/	Recycled Claim Standard Blended products included in the program use materials made from at least 50% recycled content.	f	t	\N
55	2024-02-10 19:23:13.418838+00	2024-02-10 19:23:13.418838+00	\N	Recycled Claim Standard 100	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285562/RecycledClaimStandard100.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/recycled-claim-standard-global-recycled-standard/	Recycled Claim Standard 100 products use materials made from at least 95% recycled content.	f	t	\N
56	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	Recycled Content Certification for Electrical and Electronic Equipment	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285563/RecycledContentCertificationforElectricalandElectronicEquipmentLogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.scsglobalservices.com/services/recycled-content-certification-for-electrical-and-electronic-equipment	Recycled Content Certification for Electrical and Electronic Equipment products have at least 10-50% recycled content, based on product type and environmental impact.	f	t	\N
57	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	Regenerative Organic Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285570/RegenerativeOrganicCertifiedlogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.regenorganic.org/	Regenerative Organic Certified products meet standards for soil health and land management, animal welfare, and social fairness.	f	t	\N
58	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	Responsible Wool Standard	https://res.cloudinary.com/dwmpianpg/image/upload/v1707284629/ResponsibleWoolStandard.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://textileexchange.org/standards/responsible-wool/	Responsible Wool Standard wool is from farms that support animal welfare and responsible land management practices.	f	t	\N
59	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	Soil Association	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285555/SoilAssociation.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.soilassociation.org/certification/	Soil Association certifies organic products and is a Certification Body for the COSMOS-standard, which ensures products are genuine organic or natural cosmetics.	t	t	\N
60	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	TCO Certified	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285551/TCOCertified.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://tcocertified.com/tco-certified/	TCO Certified identifies IT products that are independently assessed for lower environmental and social impact, safer chemicals, and circular design.	f	t	\N
61	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	The Forest Stewardship Council	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285567/TheForestStewardshipCouncil.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://fsc.org/en/page/about-us	Certified products support responsible forestry, helping keep forests healthy for future generations.	f	t	\N
62	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	U.S. EPA Design for the Environment	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285552/USEPADesignfortheEnvironment.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.epa.gov/pesticide-labels/learn-about-design-environment-dfe-certification	Certifies disinfectants that meet the Environmental Protection Agency's strict standards for human and environmental health.	f	t	\N
63	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	US Cotton Trust Protocol	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285571/USCottonTrustProtocol.png	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://trustuscotton.org/	U.S. Cotton Trust Protocol promotes farmers’ use of environmentally and socially responsible practices in growing more sustainable cotton.	f	t	\N
64	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	USDA Organic	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285556/USDAOrganic.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.usda.gov/topics/organic	USDA Organic products are grown and processed according to standards addressing soil and water quality, among other factors.	f	t	\N
65	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	WaterSense	https://res.cloudinary.com/dwmpianpg/image/upload/v1707285572/WaterSenseLogo.jpg	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.epa.gov/watersense	WaterSense, a voluntary label from the US Environmental Protection Agency (EPA), identifies products that meet criteria from the EPA for water-efficiency and performance.	f	t	\N
66	2024-02-10 19:25:18.93602+00	2024-02-10 19:25:18.93602+00	\N	BCorp Certified	https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Certified_B_Corporation_B_Corp_Logo_2022_Black_RGB.svg/1200px-Certified_B_Corporation_B_Corp_Logo_2022_Black_RGB.svg.png	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	https://www.bcorporation.net/en-us/certification/	description goes here	t	f	\N
\.


--
-- Data for Name: companies; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.companies (id, created_at, updated_at, deleted_at, "name json:name", url, description, user_id, is_verified, city, state, country, image_id, image) FROM stdin;
\.


--
-- Data for Name: favourites; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.favourites (id, created_at, updated_at, deleted_at, user_refer, product_id) FROM stdin;
\.


--
-- Data for Name: images; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.images (id, created_at, updated_at, deleted_at, image_location) FROM stdin;
\.


--
-- Data for Name: login_credentials; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.login_credentials (id, created_at, updated_at, deleted_at, email, password_hash, salt) FROM stdin;
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.products (id, created_at, updated_at, deleted_at, url, description, title) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, created_at, updated_at, deleted_at, first_name, last_name, login_credentials_id) FROM stdin;
\.


--
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.categories_id_seq', 1, false);


--
-- Name: certifications_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.certifications_id_seq', 66, true);


--
-- Name: companies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.companies_id_seq', 1, false);


--
-- Name: favourites_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.favourites_id_seq', 1, false);


--
-- Name: images_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.images_id_seq', 1, false);


--
-- Name: login_credentials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.login_credentials_id_seq', 1, false);


--
-- Name: products_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.products_id_seq', 1, false);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.users_id_seq', 1, false);


--
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- Name: certifications certifications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.certifications
    ADD CONSTRAINT certifications_pkey PRIMARY KEY (id);


--
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- Name: favourites favourites_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.favourites
    ADD CONSTRAINT favourites_pkey PRIMARY KEY (id);


--
-- Name: images images_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.images
    ADD CONSTRAINT images_pkey PRIMARY KEY (id);


--
-- Name: login_credentials login_credentials_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.login_credentials
    ADD CONSTRAINT login_credentials_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_categories_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_categories_deleted_at ON public.categories USING btree (deleted_at);


--
-- Name: idx_certifications_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_certifications_deleted_at ON public.certifications USING btree (deleted_at);


--
-- Name: idx_companies_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_companies_deleted_at ON public.companies USING btree (deleted_at);


--
-- Name: idx_favourites_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_favourites_deleted_at ON public.favourites USING btree (deleted_at);


--
-- Name: idx_images_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_images_deleted_at ON public.images USING btree (deleted_at);


--
-- Name: idx_login_credentials_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_login_credentials_deleted_at ON public.login_credentials USING btree (deleted_at);


--
-- Name: idx_products_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_products_deleted_at ON public.products USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: favourites fk_favourites_product; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.favourites
    ADD CONSTRAINT fk_favourites_product FOREIGN KEY (product_id) REFERENCES public.products(id);


--
-- Name: favourites fk_users_favourites; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.favourites
    ADD CONSTRAINT fk_users_favourites FOREIGN KEY (user_refer) REFERENCES public.users(id);


--
-- Name: users fk_users_login_credentials; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_login_credentials FOREIGN KEY (login_credentials_id) REFERENCES public.login_credentials(id);


--
-- Name: SCHEMA heroku_ext; Type: ACL; Schema: -; Owner: -
--

GRANT USAGE ON SCHEMA heroku_ext TO bkagnzyocfzouz;


--
-- Name: LANGUAGE plpgsql; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON LANGUAGE plpgsql TO bkagnzyocfzouz;


--
-- Name: FUNCTION pg_stat_statements_reset(userid oid, dbid oid, queryid bigint); Type: ACL; Schema: heroku_ext; Owner: -
--

GRANT ALL ON FUNCTION heroku_ext.pg_stat_statements_reset(userid oid, dbid oid, queryid bigint) TO bkagnzyocfzouz;


--
-- PostgreSQL database dump complete
--

