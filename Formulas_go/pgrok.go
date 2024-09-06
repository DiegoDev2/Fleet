package main

// Pgrok - Poor man's ngrok, multi-tenant HTTP/TCP reverse tunnel solution
// Homepage: https://github.com/pgrok/pgrok

import (
	"fmt"
	
	"os/exec"
)

func installPgrok() {
	// Método 1: Descargar y extraer .tar.gz
	pgrok_tar_url := "https://github.com/pgrok/pgrok/archive/refs/tags/v1.4.1.tar.gz"
	pgrok_cmd_tar := exec.Command("curl", "-L", pgrok_tar_url, "-o", "package.tar.gz")
	err := pgrok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgrok_zip_url := "https://github.com/pgrok/pgrok/archive/refs/tags/v1.4.1.zip"
	pgrok_cmd_zip := exec.Command("curl", "-L", pgrok_zip_url, "-o", "package.zip")
	err = pgrok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgrok_bin_url := "https://github.com/pgrok/pgrok/archive/refs/tags/v1.4.1.bin"
	pgrok_cmd_bin := exec.Command("curl", "-L", pgrok_bin_url, "-o", "binary.bin")
	err = pgrok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgrok_src_url := "https://github.com/pgrok/pgrok/archive/refs/tags/v1.4.1.src.tar.gz"
	pgrok_cmd_src := exec.Command("curl", "-L", pgrok_src_url, "-o", "source.tar.gz")
	err = pgrok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgrok_cmd_direct := exec.Command("./binary")
	err = pgrok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
