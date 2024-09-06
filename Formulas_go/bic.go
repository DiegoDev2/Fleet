package main

// Bic - C interpreter and API explorer
// Homepage: https://github.com/hexagonal-sun/bic

import (
	"fmt"
	
	"os/exec"
)

func installBic() {
	// Método 1: Descargar y extraer .tar.gz
	bic_tar_url := "https://github.com/hexagonal-sun/bic/releases/download/v1.0.0/bic-v1.0.0.tar.gz"
	bic_cmd_tar := exec.Command("curl", "-L", bic_tar_url, "-o", "package.tar.gz")
	err := bic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bic_zip_url := "https://github.com/hexagonal-sun/bic/releases/download/v1.0.0/bic-v1.0.0.zip"
	bic_cmd_zip := exec.Command("curl", "-L", bic_zip_url, "-o", "package.zip")
	err = bic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bic_bin_url := "https://github.com/hexagonal-sun/bic/releases/download/v1.0.0/bic-v1.0.0.bin"
	bic_cmd_bin := exec.Command("curl", "-L", bic_bin_url, "-o", "binary.bin")
	err = bic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bic_src_url := "https://github.com/hexagonal-sun/bic/releases/download/v1.0.0/bic-v1.0.0.src.tar.gz"
	bic_cmd_src := exec.Command("curl", "-L", bic_src_url, "-o", "source.tar.gz")
	err = bic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bic_cmd_direct := exec.Command("./binary")
	err = bic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
