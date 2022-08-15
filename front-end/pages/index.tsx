import type { NextPage } from "next";
import Head from "next/head";
import styles from "../styles/Home.module.css";
import QuizCard from "../components/QuizCard";
import { useMemo, useState } from "react";
import { shuffleArray } from "../utils/utils";
import { Kana } from "../models/models";
import { PriorityQueue, Node } from "../utils/priorityQueue";

type ScoreReport = {
  answeredQuestions: number;
  totalQuestions: number;
  correctAnswers: number;
};

type HomePageProps = {
  data: Kana[];
};

const filterKana = (
  initial: Kana[],
  key: keyof Kana,
  filter: string
): Kana[] => {
  const r = initial.filter((kana) => kana[key] === filter).map(kana=> ({...kana}));
  return r;
};

const Home: NextPage<HomePageProps> = (props: HomePageProps) => {
  const [score, setScore] = useState<ScoreReport>({
    answeredQuestions: 0,
    correctAnswers: 0,
    totalQuestions: props.data.length,
  });

  const [currentKana, setCurrentKana] = useState<Kana[]>(props.data);
  const queue = useMemo(()=> new PriorityQueue<Kana>(props.data), [props.data]);
  const [currentNode, setCurrentNode] = useState(queue.dequeue());
  
  const onAnswer = (answeredCorrectly: boolean) => {
    if (answeredCorrectly) {
      setScore((prevState) => {
        return {
          ...prevState,
          correctAnswers: prevState.correctAnswers + 1,
          answeredQuestions: prevState.answeredQuestions + 1,
        };
      });
    } else {
      setScore((prevState) => {
        return {
          ...prevState,
          answeredQuestions: prevState.answeredQuestions + 1,
        };
      });
    }
   setCurrentNode(queue.dequeue())
  };

  const resetScore = () => {
    setScore({
      answeredQuestions: 0,
      correctAnswers: 0,
      totalQuestions: currentKana.length,
    });
  };

  return (
    <div className={styles.container}>
      <Head>
        <title>Kana quiz</title>
        <meta name="description" content="Kana quiz!" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}></h1>
        <p className={styles.description}>
          Test your kana knowledge!
          {score.answeredQuestions === score.totalQuestions && (
            <div className={styles.code}>
              Your score is
              {parseFloat(
                (score.correctAnswers / score.totalQuestions).toFixed(3)
              ) * 100}
              %
            </div>
          )}
        </p>
        <section className={styles.filterSection}>
          <h2>Filter by: </h2>
          <a
            onClick={() => {
              setCurrentKana(
                filterKana(props.data, "classification", "hiragana")
              );
              resetScore();
            }}
          >
            hiragana
          </a>
          <a
            onClick={() => {
              setCurrentKana(
                filterKana(props.data, "classification", "katakana")
              );
              resetScore();
            }}
          >
            katakana
          </a>
        </section>

        <div className={styles.grid}>
          {currentKana.map((data, i) => {
            return <QuizCard key={data.classification + data.displayValue} onAnswer={onAnswer} {...data} />;
          })}
          {currentNode && <QuizCard onAnswer={onAnswer} {...currentNode.value}></QuizCard>}
        </div>
      </main>

      <footer className={styles.footer}>
        <span className={styles.logo}></span>
      </footer>
    </div>
  );
};

export async function getServerSideProps() {
  const fetchUrl = `http://${
    process.env.API_BASE_URL ?? "localhost"
  }:8000/kana`;
  const res = await fetch(fetchUrl);
  const data: Kana[] = await res.json();
  return { props: { data: shuffleArray(data) } };
}

export default Home;
