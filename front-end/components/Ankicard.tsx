import React, { useEffect, useState } from "react";
import { Box, Textarea, Text, Stack, Button } from "@chakra-ui/react";
import styles from "../styles/QuizCard.module.css";
import { FeedBackColor, QuizCardStatus } from "./types";
import { Kana } from "../models/models";
import { PriorityQueue, Node } from "../utils/priorityQueue";
interface Props {
  node: Node<Kana>;
  onAnswer: (args: any) => void;
}

const Ankicard = ({ node, onAnswer }: Props) => {
  const [answer, setAnswer] = useState<string>("");
  const [quizCardStatus, setQuizCardStatus] = useState<QuizCardStatus>("unanswered");
  const [cardColor, setCardColor] = useState<FeedBackColor>("blue.300");
  const [alreadyAnswered, setAlreadyAnswered] = useState<boolean>(false);
  const { displayValue, correctAnswer } = node.value;
  return (
    <Box
      p={4}
      alignItems="center"
      justifyContent="center"
      display='flex'
      maxWidth='40rem'
      borderWidth={1}
      margin={2}
      bg={cardColor}
      borderRadius={"15px"}
      className={quizCardStatus === "incorrectAnswer" ? styles.wrongAnswerAnimation : ""}
    >
      <Stack
        align={{ base: "center" }}
        textAlign={{ base: "center", md: "center" }}
        mt={{ base: 4, md: 0 }}
        ml={{ md: 6 }}
      >
        <Text fontWeight='bold' textTransform='uppercase' fontSize='9xl' letterSpacing='wide' color='black'>
          {displayValue}
        </Text>
        <Box bg='white' h={10} w={"10vw"}>
          {alreadyAnswered && (
            <Text fontWeight='bold' textTransform='uppercase' fontSize='xl' letterSpacing='wide' color='black'>
              {correctAnswer}
            </Text>
          )}
        </Box>
        <Stack direction='row' align='center'>
          <Button
            onClick={() => {
              setAlreadyAnswered(true);
            }}
            colorScheme='gray'
            size='md'
          >
            Show answer
          </Button>
        </Stack>
      </Stack>
    </Box>
  );
};

export default Ankicard;
