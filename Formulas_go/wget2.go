package main

// Wget2 - Successor of GNU Wget, a file and recursive website downloader
// Homepage: https://gitlab.com/gnuwget/wget2

import (
	"fmt"
	
	"os/exec"
)

func installWget2() {
	// Método 1: Descargar y extraer .tar.gz
	wget2_tar_url := "https://ftp.gnu.org/gnu/wget/wget2-2.1.0.tar.gz"
	wget2_cmd_tar := exec.Command("curl", "-L", wget2_tar_url, "-o", "package.tar.gz")
	err := wget2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wget2_zip_url := "https://ftp.gnu.org/gnu/wget/wget2-2.1.0.zip"
	wget2_cmd_zip := exec.Command("curl", "-L", wget2_zip_url, "-o", "package.zip")
	err = wget2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wget2_bin_url := "https://ftp.gnu.org/gnu/wget/wget2-2.1.0.bin"
	wget2_cmd_bin := exec.Command("curl", "-L", wget2_bin_url, "-o", "binary.bin")
	err = wget2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wget2_src_url := "https://ftp.gnu.org/gnu/wget/wget2-2.1.0.src.tar.gz"
	wget2_cmd_src := exec.Command("curl", "-L", wget2_src_url, "-o", "source.tar.gz")
	err = wget2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wget2_cmd_direct := exec.Command("./binary")
	err = wget2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: lzlib")
exec.Command("latte", "install", "lzlib")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: libpsl")
exec.Command("latte", "install", "libpsl")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
