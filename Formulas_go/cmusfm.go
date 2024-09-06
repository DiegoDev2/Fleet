package main

// Cmusfm - Last.fm standalone scrobbler for the cmus music player
// Homepage: https://github.com/Arkq/cmusfm

import (
	"fmt"
	
	"os/exec"
)

func installCmusfm() {
	// Método 1: Descargar y extraer .tar.gz
	cmusfm_tar_url := "https://github.com/Arkq/cmusfm/archive/refs/tags/v0.5.0.tar.gz"
	cmusfm_cmd_tar := exec.Command("curl", "-L", cmusfm_tar_url, "-o", "package.tar.gz")
	err := cmusfm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmusfm_zip_url := "https://github.com/Arkq/cmusfm/archive/refs/tags/v0.5.0.zip"
	cmusfm_cmd_zip := exec.Command("curl", "-L", cmusfm_zip_url, "-o", "package.zip")
	err = cmusfm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmusfm_bin_url := "https://github.com/Arkq/cmusfm/archive/refs/tags/v0.5.0.bin"
	cmusfm_cmd_bin := exec.Command("curl", "-L", cmusfm_bin_url, "-o", "binary.bin")
	err = cmusfm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmusfm_src_url := "https://github.com/Arkq/cmusfm/archive/refs/tags/v0.5.0.src.tar.gz"
	cmusfm_cmd_src := exec.Command("curl", "-L", cmusfm_src_url, "-o", "source.tar.gz")
	err = cmusfm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmusfm_cmd_direct := exec.Command("./binary")
	err = cmusfm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libfaketime")
exec.Command("latte", "install", "libfaketime")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
