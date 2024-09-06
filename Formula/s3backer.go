package main

// S3Backer - FUSE-based single file backing store via Amazon S3
// Homepage: https://github.com/archiecobbs/s3backer

import (
	"fmt"
	
	"os/exec"
)

func installS3Backer() {
	// Método 1: Descargar y extraer .tar.gz
	s3backer_tar_url := "https://s3.amazonaws.com/archie-public/s3backer/s3backer-2.1.3.tar.gz"
	s3backer_cmd_tar := exec.Command("curl", "-L", s3backer_tar_url, "-o", "package.tar.gz")
	err := s3backer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s3backer_zip_url := "https://s3.amazonaws.com/archie-public/s3backer/s3backer-2.1.3.zip"
	s3backer_cmd_zip := exec.Command("curl", "-L", s3backer_zip_url, "-o", "package.zip")
	err = s3backer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s3backer_bin_url := "https://s3.amazonaws.com/archie-public/s3backer/s3backer-2.1.3.bin"
	s3backer_cmd_bin := exec.Command("curl", "-L", s3backer_bin_url, "-o", "binary.bin")
	err = s3backer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s3backer_src_url := "https://s3.amazonaws.com/archie-public/s3backer/s3backer-2.1.3.src.tar.gz"
	s3backer_cmd_src := exec.Command("curl", "-L", s3backer_src_url, "-o", "source.tar.gz")
	err = s3backer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s3backer_cmd_direct := exec.Command("./binary")
	err = s3backer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
