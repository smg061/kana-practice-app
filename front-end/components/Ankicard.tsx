import React, { useState } from "react";
import { Box, Text, Stack, Button } from "@chakra-ui/react";
import styles from "../styles/QuizCard.module.css";
import { Kana } from "../models/models";
import { Node } from "../utils/priorityQueue";
interface Props {
  node: Node<Kana>;
  onAnswer: (args: any) => void;
}

const Ankicard = ({ node, onAnswer }: Props) => {
  const [alreadyAnswered, setAlreadyAnswered] = useState<boolean>(false);
  const { displayValue, correctAnswer } = node.value;
  return (
    <Box
      p={4}
      alignItems='center'
      justifyContent='center'
      display='flex'
      maxWidth='40rem'
      borderWidth={1}
      margin={2}
      bg={"blue.300"}
      borderRadius={"15px"}
      width={"17.5vw"}
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
          {!alreadyAnswered && (
            <Button
              onClick={() => {
                setAlreadyAnswered(true);
              }}
              colorScheme='gray'
              size='md'
            >
              Show answer
            </Button>
          )}
          {alreadyAnswered && (
            <Stack direction='row'>
              <Button onClick={() => onAnswer(10)}>Again</Button>
              <Button onClick={() => onAnswer(1)}>Good</Button>
              <Button onClick={() => onAnswer(1)}>Easy</Button>
            </Stack>
          )}
        </Stack>
      </Stack>
    </Box>
  );
};

export default Ankicard;
