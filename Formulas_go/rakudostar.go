package main

// RakudoStar - Rakudo compiler and commonly used packages
// Homepage: https://rakudo.org/

import (
	"fmt"
	
	"os/exec"
)

func installRakudoStar() {
	// Método 1: Descargar y extraer .tar.gz
	rakudostar_tar_url := "https://github.com/rakudo/star/releases/download/2024.08/rakudo-star-2024.08.tar.gz"
	rakudostar_cmd_tar := exec.Command("curl", "-L", rakudostar_tar_url, "-o", "package.tar.gz")
	err := rakudostar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rakudostar_zip_url := "https://github.com/rakudo/star/releases/download/2024.08/rakudo-star-2024.08.zip"
	rakudostar_cmd_zip := exec.Command("curl", "-L", rakudostar_zip_url, "-o", "package.zip")
	err = rakudostar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rakudostar_bin_url := "https://github.com/rakudo/star/releases/download/2024.08/rakudo-star-2024.08.bin"
	rakudostar_cmd_bin := exec.Command("curl", "-L", rakudostar_bin_url, "-o", "binary.bin")
	err = rakudostar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rakudostar_src_url := "https://github.com/rakudo/star/releases/download/2024.08/rakudo-star-2024.08.src.tar.gz"
	rakudostar_cmd_src := exec.Command("curl", "-L", rakudostar_src_url, "-o", "source.tar.gz")
	err = rakudostar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rakudostar_cmd_direct := exec.Command("./binary")
	err = rakudostar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
exec.Command("latte", "install", "bash")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
