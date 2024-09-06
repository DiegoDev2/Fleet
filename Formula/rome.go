package main

// Rome - Carthage cache for S3, Minio, Ceph, Google Storage, Artifactory and many others
// Homepage: https://github.com/tmspzz/Rome

import (
	"fmt"
	
	"os/exec"
)

func installRome() {
	// Método 1: Descargar y extraer .tar.gz
	rome_tar_url := "https://github.com/tmspzz/Rome/archive/refs/tags/v0.24.0.65.tar.gz"
	rome_cmd_tar := exec.Command("curl", "-L", rome_tar_url, "-o", "package.tar.gz")
	err := rome_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rome_zip_url := "https://github.com/tmspzz/Rome/archive/refs/tags/v0.24.0.65.zip"
	rome_cmd_zip := exec.Command("curl", "-L", rome_zip_url, "-o", "package.zip")
	err = rome_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rome_bin_url := "https://github.com/tmspzz/Rome/archive/refs/tags/v0.24.0.65.bin"
	rome_cmd_bin := exec.Command("curl", "-L", rome_bin_url, "-o", "binary.bin")
	err = rome_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rome_src_url := "https://github.com/tmspzz/Rome/archive/refs/tags/v0.24.0.65.src.tar.gz"
	rome_cmd_src := exec.Command("curl", "-L", rome_src_url, "-o", "source.tar.gz")
	err = rome_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rome_cmd_direct := exec.Command("./binary")
	err = rome_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@8.10")
	exec.Command("latte", "install", "ghc@8.10").Run()
}
