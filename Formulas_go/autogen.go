package main

// Autogen - Automated text file generator
// Homepage: https://autogen.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAutogen() {
	// Método 1: Descargar y extraer .tar.gz
	autogen_tar_url := "https://ftp.gnu.org/gnu/autogen/rel5.18.16/autogen-5.18.16.tar.xz"
	autogen_cmd_tar := exec.Command("curl", "-L", autogen_tar_url, "-o", "package.tar.gz")
	err := autogen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autogen_zip_url := "https://ftp.gnu.org/gnu/autogen/rel5.18.16/autogen-5.18.16.tar.xz"
	autogen_cmd_zip := exec.Command("curl", "-L", autogen_zip_url, "-o", "package.zip")
	err = autogen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autogen_bin_url := "https://ftp.gnu.org/gnu/autogen/rel5.18.16/autogen-5.18.16.tar.xz"
	autogen_cmd_bin := exec.Command("curl", "-L", autogen_bin_url, "-o", "binary.bin")
	err = autogen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autogen_src_url := "https://ftp.gnu.org/gnu/autogen/rel5.18.16/autogen-5.18.16.tar.xz"
	autogen_cmd_src := exec.Command("curl", "-L", autogen_src_url, "-o", "source.tar.gz")
	err = autogen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autogen_cmd_direct := exec.Command("./binary")
	err = autogen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: guile")
exec.Command("latte", "install", "guile")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
}
