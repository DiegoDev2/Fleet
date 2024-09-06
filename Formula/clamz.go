package main

// Clamz - Download MP3 files from Amazon's music store
// Homepage: https://code.google.com/archive/p/clamz/

import (
	"fmt"
	
	"os/exec"
)

func installClamz() {
	// Método 1: Descargar y extraer .tar.gz
	clamz_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/clamz/clamz-0.5.tar.gz"
	clamz_cmd_tar := exec.Command("curl", "-L", clamz_tar_url, "-o", "package.tar.gz")
	err := clamz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clamz_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/clamz/clamz-0.5.zip"
	clamz_cmd_zip := exec.Command("curl", "-L", clamz_zip_url, "-o", "package.zip")
	err = clamz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clamz_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/clamz/clamz-0.5.bin"
	clamz_cmd_bin := exec.Command("curl", "-L", clamz_bin_url, "-o", "binary.bin")
	err = clamz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clamz_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/clamz/clamz-0.5.src.tar.gz"
	clamz_cmd_src := exec.Command("curl", "-L", clamz_src_url, "-o", "source.tar.gz")
	err = clamz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clamz_cmd_direct := exec.Command("./binary")
	err = clamz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
