import { Text } from "@chakra-ui/react";

import { Hero } from "../components/Hero";
import { Box } from "@chakra-ui/react";

import { Container } from "../components/Container";
import { Main } from "../components/Main";
import { DarkModeSwitch } from "../components/DarkModeSwitch";
import { Footer } from "../components/Footer";

const Index = () => (
  <Container height="100vh">
    <Hero title="This is a welcome title" />
    <Main>
      <Box display="flex" justifyContent="center">
        Hello sir
      </Box>
    </Main>

    <DarkModeSwitch />

    <Footer>
      <Text>This is a chakra footer</Text>
    </Footer>
  </Container>
);

export default Index;
