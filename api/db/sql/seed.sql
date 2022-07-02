--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Debian 14.3-1.pgdg110+1)
-- Dumped by pg_dump version 14.3 (Debian 14.3-1.pgdg110+1)

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
-- Data for Name: kana; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (87, 'あ', 'a', 1, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (88, 'い', 'i', 1, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (89, 'う', 'u', 1, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (90, 'え', 'e', 1, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (91, 'お', 'o', 1, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (92, 'か', 'ka', 1, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (93, 'き', 'ki', 1, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (94, 'く', 'ku', 1, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (95, 'け', 'ke', 1, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (96, 'こ', 'ko', 1, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (97, 'ま', 'ma', 1, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (98, 'み', 'mi', 1, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (99, 'む', 'mu', 1, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (100, 'め', 'me', 1, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (101, 'も', 'mo', 1, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (102, 'な', 'na', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (103, 'に', 'ni', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (104, 'ぬ', 'nu', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (105, 'ね', 'ne', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (106, 'の', 'no', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (107, 'は', 'ha', 1, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (108, 'へ', 'he', 1, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (109, 'ふ', 'fu', 1, 'f');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (110, 'ひ', 'hi', 1, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (111, 'ほ', 'ho', 1, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (112, 'た', 'ta', 1, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (113, 'ち', 'chi', 1, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (114, 'つ', 'tsu', 1, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (115, 'て', 'te', 1, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (116, 'と', 'to', 1, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (117, 'や', 'ya', 1, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (118, 'よ', 'yo', 1, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (119, 'ゆ', 'yu', 1, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (120, 'ら', 'ra', 1, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (121, 'れ', 're', 1, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (122, 'る', 'ru', 1, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (123, 'り', 'ri', 1, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (124, 'ろ', 'ro', 1, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (125, 'さ', 'sa', 1, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (126, 'し', 'shi', 1, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (127, 'す', 'su', 1, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (128, 'せ', 'se', 1, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (129, 'そ', 'so', 1, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (130, 'わ', 'wa', 1, 'w');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (131, 'を', 'wo', 1, 'w');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (132, 'ん', 'n', 1, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (133, 'ア', 'a', 2, 'a');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (134, 'イ', 'i', 2, 'i');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (135, 'ウ', 'u', 2, 'u');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (136, 'エ', 'e', 2, 'e');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (137, 'オ', 'o', 2, 'o');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (138, 'カ', 'ka', 2, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (139, 'キ', 'ki', 2, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (140, 'ク', 'ku', 2, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (141, 'ケ', 'ke', 2, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (142, 'コ', 'ko', 2, 'k');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (143, 'マ', 'ma', 2, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (144, 'ミ', 'mi', 2, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (145, 'ム', 'mu', 2, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (146, 'メ', 'me', 2, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (147, 'モ', 'mo', 2, 'm');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (148, 'ナ', 'na', 2, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (149, 'ニ', 'ni', 2, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (150, 'ヌ', 'nu', 2, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (151, 'ネ', 'ne', 2, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (152, 'ノ', 'no', 2, 'n');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (153, 'ハ', 'ha', 2, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (154, 'ヘ', 'he', 2, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (155, 'フ', 'fu', 2, 'f');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (156, 'ヒ', 'hi', 2, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (157, 'ホ', 'ho', 2, 'h');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (158, 'タ', 'ta', 2, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (159, 'チ', 'chi', 2, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (160, 'ツ', 'tsu', 2, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (161, 'テ', 'te', 2, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (162, 'ト', 'to', 2, 't');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (163, 'ヨ', 'yo', 2, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (164, 'ヤ', 'ya', 2, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (165, 'ユ', 'yu', 2, 'y');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (166, 'ラ', 'ra', 2, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (167, 'レ', 're', 2, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (168, 'ル', 'ru', 2, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (169, 'リ', 'ri', 2, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (170, 'ロ', 'ro', 2, 'r');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (171, 'サ', 'sa', 2, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (172, 'シ', 'shi', 2, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (173, 'ス', 'su', 2, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (174, 'セ', 'se', 2, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (175, 'ソ', 'so', 2, 's');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (176, 'ワ', 'wa', 2, 'w');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (177, 'ヲ', 'wo', 2, 'w');
INSERT INTO public.kana (id, representation, romaji, classification, initial) VALUES (178, 'ン', 'n', 2, 'n');

