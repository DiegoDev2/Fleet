package main

// Mdbtools - Tools to facilitate the use of Microsoft Access databases
// Homepage: https://github.com/mdbtools/mdbtools/

import (
	"fmt"
	
	"os/exec"
)

func installMdbtools() {
	// Método 1: Descargar y extraer .tar.gz
	mdbtools_tar_url := "https://github.com/mdbtools/mdbtools/releases/download/v1.0.0/mdbtools-1.0.0.tar.gz"
	mdbtools_cmd_tar := exec.Command("curl", "-L", mdbtools_tar_url, "-o", "package.tar.gz")
	err := mdbtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdbtools_zip_url := "https://github.com/mdbtools/mdbtools/releases/download/v1.0.0/mdbtools-1.0.0.zip"
	mdbtools_cmd_zip := exec.Command("curl", "-L", mdbtools_zip_url, "-o", "package.zip")
	err = mdbtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdbtools_bin_url := "https://github.com/mdbtools/mdbtools/releases/download/v1.0.0/mdbtools-1.0.0.bin"
	mdbtools_cmd_bin := exec.Command("curl", "-L", mdbtools_bin_url, "-o", "binary.bin")
	err = mdbtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdbtools_src_url := "https://github.com/mdbtools/mdbtools/releases/download/v1.0.0/mdbtools-1.0.0.src.tar.gz"
	mdbtools_cmd_src := exec.Command("curl", "-L", mdbtools_src_url, "-o", "source.tar.gz")
	err = mdbtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdbtools_cmd_direct := exec.Command("./binary")
	err = mdbtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
