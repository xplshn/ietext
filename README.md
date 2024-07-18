# A text editor that comes in different parts
- ietextl will add line numbers (like the C64's)
- ietextil will remove the line numbers, thus converting the text to a normal file when using a pipe
- ietextw provides a text editing window. Where CTRL+Q is copy, CTRL+P is paste, CTRL+X is cut and CTRL+Z is undo. Shift+ArrowKey is select and so on.

#### The idea
You can basically override lines by specifiying their line numbers. Same way an unstructured text editor would work. For example:
`ietextl <~/.shrc | ietextw | ietextil > ./.shrc.2`
![2024-07-18-163051_1366x768_scrot](https://github.com/user-attachments/assets/7bd85be3-c56a-4859-ba63-94c998802cbd)

> [!IMPORTANT]
> Currently, whitespace is destroyed by ietextw's output :(
> I haven't had the time to correct the issue/bug. And given that this serves no purpose. I probably won't be fixing it any time soon
