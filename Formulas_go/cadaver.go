package main

// Cadaver - Command-line client for DAV
// Homepage: https://notroj.github.io/cadaver/

import (
	"fmt"
	
	"os/exec"
)

func installCadaver() {
	// Método 1: Descargar y extraer .tar.gz
	cadaver_tar_url := "https://notroj.github.io/cadaver/cadaver-0.24.tar.gz"
	cadaver_cmd_tar := exec.Command("curl", "-L", cadaver_tar_url, "-o", "package.tar.gz")
	err := cadaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cadaver_zip_url := "https://notroj.github.io/cadaver/cadaver-0.24.zip"
	cadaver_cmd_zip := exec.Command("curl", "-L", cadaver_zip_url, "-o", "package.zip")
	err = cadaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cadaver_bin_url := "https://notroj.github.io/cadaver/cadaver-0.24.bin"
	cadaver_cmd_bin := exec.Command("curl", "-L", cadaver_bin_url, "-o", "binary.bin")
	err = cadaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cadaver_src_url := "https://notroj.github.io/cadaver/cadaver-0.24.src.tar.gz"
	cadaver_cmd_src := exec.Command("curl", "-L", cadaver_src_url, "-o", "source.tar.gz")
	err = cadaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cadaver_cmd_direct := exec.Command("./binary")
	err = cadaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: neon")
exec.Command("latte", "install", "neon")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
