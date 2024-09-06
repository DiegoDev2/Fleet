package main

// Seqtk - Toolkit for processing sequences in FASTA/Q formats
// Homepage: https://github.com/lh3/seqtk

import (
	"fmt"
	
	"os/exec"
)

func installSeqtk() {
	// Método 1: Descargar y extraer .tar.gz
	seqtk_tar_url := "https://github.com/lh3/seqtk/archive/refs/tags/v1.4.tar.gz"
	seqtk_cmd_tar := exec.Command("curl", "-L", seqtk_tar_url, "-o", "package.tar.gz")
	err := seqtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seqtk_zip_url := "https://github.com/lh3/seqtk/archive/refs/tags/v1.4.zip"
	seqtk_cmd_zip := exec.Command("curl", "-L", seqtk_zip_url, "-o", "package.zip")
	err = seqtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seqtk_bin_url := "https://github.com/lh3/seqtk/archive/refs/tags/v1.4.bin"
	seqtk_cmd_bin := exec.Command("curl", "-L", seqtk_bin_url, "-o", "binary.bin")
	err = seqtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seqtk_src_url := "https://github.com/lh3/seqtk/archive/refs/tags/v1.4.src.tar.gz"
	seqtk_cmd_src := exec.Command("curl", "-L", seqtk_src_url, "-o", "source.tar.gz")
	err = seqtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seqtk_cmd_direct := exec.Command("./binary")
	err = seqtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
