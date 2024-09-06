package main

// Gengetopt - Generate C code to parse command-line arguments via getopt_long
// Homepage: https://www.gnu.org/software/gengetopt/

import (
	"fmt"
	
	"os/exec"
)

func installGengetopt() {
	// Método 1: Descargar y extraer .tar.gz
	gengetopt_tar_url := "https://ftp.gnu.org/gnu/gengetopt/gengetopt-2.23.tar.xz"
	gengetopt_cmd_tar := exec.Command("curl", "-L", gengetopt_tar_url, "-o", "package.tar.gz")
	err := gengetopt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gengetopt_zip_url := "https://ftp.gnu.org/gnu/gengetopt/gengetopt-2.23.tar.xz"
	gengetopt_cmd_zip := exec.Command("curl", "-L", gengetopt_zip_url, "-o", "package.zip")
	err = gengetopt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gengetopt_bin_url := "https://ftp.gnu.org/gnu/gengetopt/gengetopt-2.23.tar.xz"
	gengetopt_cmd_bin := exec.Command("curl", "-L", gengetopt_bin_url, "-o", "binary.bin")
	err = gengetopt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gengetopt_src_url := "https://ftp.gnu.org/gnu/gengetopt/gengetopt-2.23.tar.xz"
	gengetopt_cmd_src := exec.Command("curl", "-L", gengetopt_src_url, "-o", "source.tar.gz")
	err = gengetopt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gengetopt_cmd_direct := exec.Command("./binary")
	err = gengetopt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
