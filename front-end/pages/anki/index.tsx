import type { NextPage } from "next";
import Head from "next/head";
import styles from "../../styles/Home.module.css";
import { useEffect, useMemo, useState } from "react";
import { shuffleArray } from "../../utils/utils";
import { Kana } from "../../models/models";
import { PriorityQueue, Node } from "../../utils/priorityQueue";
import Ankicard from "../../components/Ankicard";

type ScoreReport = {
  answeredQuestions: number;
  totalQuestions: number;
  correctAnswers: number;
};

type HomePageProps = {
  data: Kana[];
};

const filterKana = (initial: Kana[], key: keyof Kana, filter: string): Kana[] => {
  const r = initial.filter((kana) => kana[key] === filter).map((kana) => ({ ...kana }));
  return r;
};

const Home: NextPage<HomePageProps> = ({data}: HomePageProps) => {
  const [score, setScore] = useState<ScoreReport>({
    answeredQuestions: 0,
    correctAnswers: 0,
    totalQuestions: data.length,
  });

  const [currentKana, setCurrentKana] = useState<Kana[]>(data);
  const queue = useMemo(() => new PriorityQueue<Kana>(currentKana), [currentKana]);
  const [currentNode, setCurrentNode] = useState(queue.dequeue());

  const onAnswer = (args: any) => {
    setCurrentNode(queue.dequeue());
  };

  const resetScore = () => {
    setScore({
      answeredQuestions: 0,
      correctAnswers: 0,
      totalQuestions: currentKana.length,
    });
  };

  useEffect(()=> {
    setCurrentNode(queue.dequeue())
  }, [2queue])

  return (
    <div className={styles.container}>
      <Head>
        <title>Kana quiz</title>
        <meta name='description' content='Kana quiz!' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}></h1>
        <p className={styles.description}>
          Test your kana knowledge!
          {score.answeredQuestions === score.totalQuestions && (
            <div className={styles.code}>
              Your score is
              {parseFloat((score.correctAnswers / score.totalQuestions).toFixed(3)) * 100}%
            </div>
          )}
        </p>
        <section className={styles.filterSection}>
          <h2>Filter by: </h2>
          <a
            onClick={() => {
              const newKana = filterKana(data, "classification", "hiragana")
              setCurrentKana(newKana);
              resetScore();
            }}
          >
            hiragana
          </a>
          <a
            onClick={() => {
              const newKana = filterKana(data, "classification", "katakana")
              setCurrentKana(newKana);
              resetScore();
            }}
          >
            katakana
          </a>
        </section>

        <div className={styles.grid}>
          {currentNode && (
            <Ankicard
              onAnswer={onAnswer}
              key={currentNode.value.classification + currentNode.value.correctAnswer}
              node={currentNode}
            ></Ankicard>
          )}
        </div>
      </main>

      <footer className={styles.footer}>
        <span className={styles.logo}></span>
      </footer>
    </div>
  );
};

export async function getServerSideProps() {
  const fetchUrl = `http://${process.env.API_BASE_URL ?? "localhost"}:8000/kana`;
  const res = await fetch(fetchUrl);
  const data: Kana[] = await res.json();
  return { props: { data: shuffleArray(data) } };
}

export default Home;
