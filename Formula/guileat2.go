package main

// GuileAT2 - GNU Ubiquitous Intelligent Language for Extensions
// Homepage: https://www.gnu.org/software/guile/

import (
	"fmt"
	
	"os/exec"
)

func installGuileAT2() {
	// Método 1: Descargar y extraer .tar.gz
	guileat2_tar_url := "https://ftp.gnu.org/gnu/guile/guile-2.2.7.tar.xz"
	guileat2_cmd_tar := exec.Command("curl", "-L", guileat2_tar_url, "-o", "package.tar.gz")
	err := guileat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	guileat2_zip_url := "https://ftp.gnu.org/gnu/guile/guile-2.2.7.tar.xz"
	guileat2_cmd_zip := exec.Command("curl", "-L", guileat2_zip_url, "-o", "package.zip")
	err = guileat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	guileat2_bin_url := "https://ftp.gnu.org/gnu/guile/guile-2.2.7.tar.xz"
	guileat2_cmd_bin := exec.Command("curl", "-L", guileat2_bin_url, "-o", "binary.bin")
	err = guileat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	guileat2_src_url := "https://ftp.gnu.org/gnu/guile/guile-2.2.7.tar.xz"
	guileat2_cmd_src := exec.Command("curl", "-L", guileat2_src_url, "-o", "source.tar.gz")
	err = guileat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	guileat2_cmd_direct := exec.Command("./binary")
	err = guileat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libunistring")
	exec.Command("latte", "install", "libunistring").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
