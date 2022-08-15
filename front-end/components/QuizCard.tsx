import React, { useEffect, useState } from "react";
import { Box, Textarea, Text, Stack } from "@chakra-ui/react";
import styles from "../styles/QuizCard.module.css";
interface Props {
  displayValue: string;
  correctAnswer: string;
  onAnswer: (answeeredCorrectly: boolean) => void;
}

type QuizCardStatus = "unanswered" | "correctAnswer" | "incorrectAnswer";
type FeedBackColor = "tomato" | "green.300" | "blue.300";

const QuizCard = ({ displayValue, correctAnswer, onAnswer }: Props) => {
  const [answer, setAnswer] = useState<string>("");
  const [quizCardStatus, setQuizCardStatus] =
    useState<QuizCardStatus>("unanswered");
  const [cardColor, setCardColor] = useState<FeedBackColor>("blue.300");
  const [alreadyAnswered, setAlreadyAnswered] = useState<boolean>(false);

  const handleAnswer = (value: string) => {
    if (value === "\n") return;
    setAnswer(value);
  };

  useEffect(()=> {
    setCardColor('blue.300');
    setQuizCardStatus('unanswered');
    setAlreadyAnswered(false);
  },[displayValue,correctAnswer])
  return (
      <Box
        p={4}
        display={{ md: "flex" }}
        maxWidth="32rem"
        borderWidth={1}
        margin={2}
        bg={cardColor}
        borderRadius={"15px"}
        className={
          quizCardStatus === "incorrectAnswer"
            ? styles.wrongAnswerAnimation
            : ""
        }
      >
        <Stack
          align={{ base: "center", md: "stretch" }}
          textAlign={{ base: "center", md: "left" }}
          mt={{ base: 4, md: 0 }}
          ml={{ md: 6 }}
        >
          <Text
            fontWeight="bold"
            textTransform="uppercase"
            fontSize="6xl"
            letterSpacing="wide"
            color="black"
          >
            {displayValue}
          </Text>
          <Textarea
            value={answer}
            onChange={(e) => handleAnswer(e.target.value)}
            disabled={alreadyAnswered}
            bg="white"
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                if (!answer) return;
                if (answer.toLocaleLowerCase() === correctAnswer) {
                  setQuizCardStatus("correctAnswer");
                  setCardColor("green.300");
                  onAnswer(true);
                } else {
                  setQuizCardStatus("incorrectAnswer");
                  setCardColor("tomato");
                  onAnswer(false);
                }
                setAlreadyAnswered(true);
              }
            }}
            placeholder="Enter your answer"
            resize="none"
          ></Textarea>
        </Stack>
      </Box>
  );
};

export default QuizCard;
