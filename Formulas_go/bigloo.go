package main

// Bigloo - Scheme implementation with object system, C, and Java interfaces
// Homepage: https://www-sop.inria.fr/indes/fp/Bigloo/

import (
	"fmt"
	
	"os/exec"
)

func installBigloo() {
	// Método 1: Descargar y extraer .tar.gz
	bigloo_tar_url := "ftp://ftp-sop.inria.fr/indes/fp/Bigloo/bigloo-4.5b.tar.gz"
	bigloo_cmd_tar := exec.Command("curl", "-L", bigloo_tar_url, "-o", "package.tar.gz")
	err := bigloo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bigloo_zip_url := "ftp://ftp-sop.inria.fr/indes/fp/Bigloo/bigloo-4.5b.zip"
	bigloo_cmd_zip := exec.Command("curl", "-L", bigloo_zip_url, "-o", "package.zip")
	err = bigloo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bigloo_bin_url := "ftp://ftp-sop.inria.fr/indes/fp/Bigloo/bigloo-4.5b.bin"
	bigloo_cmd_bin := exec.Command("curl", "-L", bigloo_bin_url, "-o", "binary.bin")
	err = bigloo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bigloo_src_url := "ftp://ftp-sop.inria.fr/indes/fp/Bigloo/bigloo-4.5b.src.tar.gz"
	bigloo_cmd_src := exec.Command("curl", "-L", bigloo_src_url, "-o", "source.tar.gz")
	err = bigloo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bigloo_cmd_direct := exec.Command("./binary")
	err = bigloo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
