# A text editor that comes in different parts
- ietextl will add line numbers (like the C64's)
- ietextil will remove the line numbers, thus converting the text to a normal file when using a pipe
- ietextw provides a text editing window. Where CTRL+Q is copy, CTRL+P is paste, CTRL+X is cut and CTRL+Z is undo. Shift+ArrowKey is select and so on.

#### The idea
You can basically override lines by specifiying their line numbers. Same way an unstructured text editor would work. For example:
`ietextl <~/.shrc | ietextw | ietextil > ./.shrc.2`
![2024-07-18-163051_1366x768_scrot](https://github.com/user-attachments/assets/7bd85be3-c56a-4859-ba63-94c998802cbd)

> [!IMPORTANT]
> Currently, whitespace is destroyed by ietextil's output :(
> I haven't had the time to correct the issue/bug. And given that this serves no purpose. I probably won't be fixing it any time soon, since this is only a proof of concept.

> [!NOTE]
> I've been working on adding syntax highlight to the `Ed` (the Standard Unix Text EDitor) implementation made by [u-root](https://github.com/u-root/u-root). Which can be found here: https://github.com/u-root/u-root/tree/main/cmds/exp/ed
> This is probably more useful. And if I get it working I'll replace ietextw with an enhanced `Ed`
