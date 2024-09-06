package main

// CrunchyCli - Command-line downloader for Crunchyroll
// Homepage: https://github.com/crunchy-labs/crunchy-cli

import (
	"fmt"
	
	"os/exec"
)

func installCrunchyCli() {
	// Método 1: Descargar y extraer .tar.gz
	crunchycli_tar_url := "https://github.com/crunchy-labs/crunchy-cli/archive/refs/tags/v3.6.7.tar.gz"
	crunchycli_cmd_tar := exec.Command("curl", "-L", crunchycli_tar_url, "-o", "package.tar.gz")
	err := crunchycli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crunchycli_zip_url := "https://github.com/crunchy-labs/crunchy-cli/archive/refs/tags/v3.6.7.zip"
	crunchycli_cmd_zip := exec.Command("curl", "-L", crunchycli_zip_url, "-o", "package.zip")
	err = crunchycli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crunchycli_bin_url := "https://github.com/crunchy-labs/crunchy-cli/archive/refs/tags/v3.6.7.bin"
	crunchycli_cmd_bin := exec.Command("curl", "-L", crunchycli_bin_url, "-o", "binary.bin")
	err = crunchycli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crunchycli_src_url := "https://github.com/crunchy-labs/crunchy-cli/archive/refs/tags/v3.6.7.src.tar.gz"
	crunchycli_cmd_src := exec.Command("curl", "-L", crunchycli_src_url, "-o", "source.tar.gz")
	err = crunchycli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crunchycli_cmd_direct := exec.Command("./binary")
	err = crunchycli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
