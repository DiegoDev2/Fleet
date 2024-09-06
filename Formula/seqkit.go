package main

// Seqkit - Cross-platform and ultrafast toolkit for FASTA/Q file manipulation in Golang
// Homepage: https://bioinf.shenwei.me/seqkit

import (
	"fmt"
	
	"os/exec"
)

func installSeqkit() {
	// Método 1: Descargar y extraer .tar.gz
	seqkit_tar_url := "https://github.com/shenwei356/seqkit/archive/refs/tags/v2.8.2.tar.gz"
	seqkit_cmd_tar := exec.Command("curl", "-L", seqkit_tar_url, "-o", "package.tar.gz")
	err := seqkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seqkit_zip_url := "https://github.com/shenwei356/seqkit/archive/refs/tags/v2.8.2.zip"
	seqkit_cmd_zip := exec.Command("curl", "-L", seqkit_zip_url, "-o", "package.zip")
	err = seqkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seqkit_bin_url := "https://github.com/shenwei356/seqkit/archive/refs/tags/v2.8.2.bin"
	seqkit_cmd_bin := exec.Command("curl", "-L", seqkit_bin_url, "-o", "binary.bin")
	err = seqkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seqkit_src_url := "https://github.com/shenwei356/seqkit/archive/refs/tags/v2.8.2.src.tar.gz"
	seqkit_cmd_src := exec.Command("curl", "-L", seqkit_src_url, "-o", "source.tar.gz")
	err = seqkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seqkit_cmd_direct := exec.Command("./binary")
	err = seqkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
