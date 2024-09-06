package main

// Mp3unicode - Command-line utility to convert mp3 tags between different encodings
// Homepage: https://mp3unicode.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMp3unicode() {
	// Método 1: Descargar y extraer .tar.gz
	mp3unicode_tar_url := "https://github.com/alonbl/mp3unicode/releases/download/mp3unicode-1.2.1/mp3unicode-1.2.1.tar.bz2"
	mp3unicode_cmd_tar := exec.Command("curl", "-L", mp3unicode_tar_url, "-o", "package.tar.gz")
	err := mp3unicode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3unicode_zip_url := "https://github.com/alonbl/mp3unicode/releases/download/mp3unicode-1.2.1/mp3unicode-1.2.1.tar.bz2"
	mp3unicode_cmd_zip := exec.Command("curl", "-L", mp3unicode_zip_url, "-o", "package.zip")
	err = mp3unicode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3unicode_bin_url := "https://github.com/alonbl/mp3unicode/releases/download/mp3unicode-1.2.1/mp3unicode-1.2.1.tar.bz2"
	mp3unicode_cmd_bin := exec.Command("curl", "-L", mp3unicode_bin_url, "-o", "binary.bin")
	err = mp3unicode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3unicode_src_url := "https://github.com/alonbl/mp3unicode/releases/download/mp3unicode-1.2.1/mp3unicode-1.2.1.tar.bz2"
	mp3unicode_cmd_src := exec.Command("curl", "-L", mp3unicode_src_url, "-o", "source.tar.gz")
	err = mp3unicode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3unicode_cmd_direct := exec.Command("./binary")
	err = mp3unicode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
}
