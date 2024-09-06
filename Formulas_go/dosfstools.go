package main

// Dosfstools - Tools to create, check and label file systems of the FAT family
// Homepage: https://github.com/dosfstools

import (
	"fmt"
	
	"os/exec"
)

func installDosfstools() {
	// Método 1: Descargar y extraer .tar.gz
	dosfstools_tar_url := "https://github.com/dosfstools/dosfstools/releases/download/v4.2/dosfstools-4.2.tar.gz"
	dosfstools_cmd_tar := exec.Command("curl", "-L", dosfstools_tar_url, "-o", "package.tar.gz")
	err := dosfstools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dosfstools_zip_url := "https://github.com/dosfstools/dosfstools/releases/download/v4.2/dosfstools-4.2.zip"
	dosfstools_cmd_zip := exec.Command("curl", "-L", dosfstools_zip_url, "-o", "package.zip")
	err = dosfstools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dosfstools_bin_url := "https://github.com/dosfstools/dosfstools/releases/download/v4.2/dosfstools-4.2.bin"
	dosfstools_cmd_bin := exec.Command("curl", "-L", dosfstools_bin_url, "-o", "binary.bin")
	err = dosfstools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dosfstools_src_url := "https://github.com/dosfstools/dosfstools/releases/download/v4.2/dosfstools-4.2.src.tar.gz"
	dosfstools_cmd_src := exec.Command("curl", "-L", dosfstools_src_url, "-o", "source.tar.gz")
	err = dosfstools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dosfstools_cmd_direct := exec.Command("./binary")
	err = dosfstools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
