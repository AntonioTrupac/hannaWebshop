import { Flex, Heading } from "@chakra-ui/react";

type HeroProps = {
  title?: string;
};

export const Hero = ({ title }: HeroProps) => (
  <Flex
    justifyContent="center"
    alignItems="center"
    height="100%"
    width="100%"
    bgGradient="linear(to-l, #7928CA, #FF0080)"
    bgClip="text"
  >
    <Heading fontSize="6vw">{title}</Heading>
  </Flex>
);
