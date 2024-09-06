package main

// GnuComplexity - Measures complexity of C source
// Homepage: https://www.gnu.org/software/complexity/

import (
	"fmt"
	
	"os/exec"
)

func installGnuComplexity() {
	// Método 1: Descargar y extraer .tar.gz
	gnucomplexity_tar_url := "https://ftp.gnu.org/gnu/complexity/complexity-1.13.tar.xz"
	gnucomplexity_cmd_tar := exec.Command("curl", "-L", gnucomplexity_tar_url, "-o", "package.tar.gz")
	err := gnucomplexity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnucomplexity_zip_url := "https://ftp.gnu.org/gnu/complexity/complexity-1.13.tar.xz"
	gnucomplexity_cmd_zip := exec.Command("curl", "-L", gnucomplexity_zip_url, "-o", "package.zip")
	err = gnucomplexity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnucomplexity_bin_url := "https://ftp.gnu.org/gnu/complexity/complexity-1.13.tar.xz"
	gnucomplexity_cmd_bin := exec.Command("curl", "-L", gnucomplexity_bin_url, "-o", "binary.bin")
	err = gnucomplexity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnucomplexity_src_url := "https://ftp.gnu.org/gnu/complexity/complexity-1.13.tar.xz"
	gnucomplexity_cmd_src := exec.Command("curl", "-L", gnucomplexity_src_url, "-o", "source.tar.gz")
	err = gnucomplexity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnucomplexity_cmd_direct := exec.Command("./binary")
	err = gnucomplexity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: autogen")
exec.Command("latte", "install", "autogen")
}
