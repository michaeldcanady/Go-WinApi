<h1 align="center">
  <br>
  <a href="http://www.amitmerchant.com/electron-markdownify"><img src="../../../Images/Go-WinApi_Logo.png" alt="Markdownify" width="200"></a>
  <br>
  Go-FileApi
  <br>
</h1>

<h4 align="center">A Golang Wrapper For The File  Api within the Windows Api.</h4>

<p align="center">
  <a href="https://badge.fury.io/gh/michaeldcanady%2FGo-WinApi"><img src="https://badge.fury.io/gh/michaeldcanady%2FGo-WinApi.svg" alt="GitHub version" height="18"></a>
</p>

<p align="center">
  <a href="#what-is-go-file-api">What is Go-FileApi?</a> •
      <a href="#functions">Functions</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#credits">Credits</a> •
  <a href="#related">Related</a> •
  <a href="#license">License</a>
</p>

![screenshot](https://raw.githubusercontent.com/amitmerchant1990/electron-markdownify/master/app/img/markdownify.gif)

## What is Go File Api?

Go-WinApi is a golang wrapper for the windows api. Unlike other wrappers though it parses the data provided by the API into human readable and immediately in application.

## Functions

<details>
  <summary>AreFileApisANSI(): Returns Bool</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
        <BLOCKQUOTE>
          If the set of file I/O functions is using the ANSI code page, the return value is nonzero.
          <br>
          If the set of file I/O functions is using the OEM code page, the return value is zero.
        </BLOCKQUOTE>
    </details>
    <details>
      <summary>Example</summary>
      <pre><code>
      package main
      <br>
      import(
        "fmt"
        fileapi "github.com/michaeldcanady/Go-WinApi/Go-WinApi/windows-api/kernel32/FileApi"
      )
      <br>
      func main() {
        ANSI := fileapi.AreFileApisANSI()
        fmt.Println(ANSI)
        }
        </code></pre>
    </details>
  </BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>CreateFile2(): Returns syscall.Handle, error</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>CreateFileW(): Returns syscall.Handle, error</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>DeleteFileW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindFirstFileExW(): Returns syscall.Handle, Win32FindDataW, error</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindFirstFileNameW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindFirstFileNameW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindNexFileW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindNextVolume(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>FindVolumeClose(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetDriveType(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetFixedDriveMounts(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetLogicalDrives(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetVolumeInformationW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetVolumeNameFromVolumeMountPointW(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

<details>
  <summary>GetVolumePathNameForVolumeName(): Returns</summary>
  <BLOCKQUOTE>
    <details>
      <summary>Description</summary>
      <br>
    </details>
    <details>
      <summary>Example</summary>
      <br>
    </details>
  <BLOCKQUOTE>
</details>

<hr>

## How to use

Install a ()[]

```sh
git clone https://github.com/michaeldcanady/Go-WinApi.git
```

## Contributing

## Credit

## Related

## License

N/A

---

<!-- >> [amitmerchant.com](https://www.amitmerchant.com) &nbsp;&middot;&nbsp;
> GitHub [@amitmerchant1990](https://github.com/amitmerchant1990) &nbsp;&middot;&nbsp;
> Twitter [@amit_merchant](https://twitter.com/amit_merchant) -->