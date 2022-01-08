import { Flex, FlexProps } from "@chakra-ui/react";

export const Footer = (props: FlexProps) => (
  <Flex as="footer" py="2rem" position="fixed" bottom="0" {...props} />
);
