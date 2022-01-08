import {
  Link as ChakraLink,
  Text,
  Code,
  List,
  ListIcon,
  ListItem,
} from "@chakra-ui/react";
import { CheckCircleIcon, LinkIcon } from "@chakra-ui/icons";

import { Hero } from "../components/Hero";
import { Box } from "@chakra-ui/react";

import { Container } from "../components/Container";
import { Main } from "../components/Main";
import { DarkModeSwitch } from "../components/DarkModeSwitch";
import { CTA } from "../components/CTA";
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
