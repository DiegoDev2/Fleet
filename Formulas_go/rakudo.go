package main

// Rakudo - Mature, production-ready implementation of the Raku language
// Homepage: https://rakudo.org

import (
	"fmt"
	
	"os/exec"
)

func installRakudo() {
	// Método 1: Descargar y extraer .tar.gz
	rakudo_tar_url := "https://github.com/rakudo/rakudo/releases/download/2024.08/rakudo-2024.08.tar.gz"
	rakudo_cmd_tar := exec.Command("curl", "-L", rakudo_tar_url, "-o", "package.tar.gz")
	err := rakudo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rakudo_zip_url := "https://github.com/rakudo/rakudo/releases/download/2024.08/rakudo-2024.08.zip"
	rakudo_cmd_zip := exec.Command("curl", "-L", rakudo_zip_url, "-o", "package.zip")
	err = rakudo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rakudo_bin_url := "https://github.com/rakudo/rakudo/releases/download/2024.08/rakudo-2024.08.bin"
	rakudo_cmd_bin := exec.Command("curl", "-L", rakudo_bin_url, "-o", "binary.bin")
	err = rakudo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rakudo_src_url := "https://github.com/rakudo/rakudo/releases/download/2024.08/rakudo-2024.08.src.tar.gz"
	rakudo_cmd_src := exec.Command("curl", "-L", rakudo_src_url, "-o", "source.tar.gz")
	err = rakudo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rakudo_cmd_direct := exec.Command("./binary")
	err = rakudo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtommath")
exec.Command("latte", "install", "libtommath")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: moarvm")
exec.Command("latte", "install", "moarvm")
	fmt.Println("Instalando dependencia: nqp")
exec.Command("latte", "install", "nqp")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
