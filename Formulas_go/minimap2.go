package main

// Minimap2 - Versatile pairwise aligner for genomic and spliced nucleotide sequences
// Homepage: https://lh3.github.io/minimap2

import (
	"fmt"
	
	"os/exec"
)

func installMinimap2() {
	// Método 1: Descargar y extraer .tar.gz
	minimap2_tar_url := "https://github.com/lh3/minimap2/archive/refs/tags/v2.28.tar.gz"
	minimap2_cmd_tar := exec.Command("curl", "-L", minimap2_tar_url, "-o", "package.tar.gz")
	err := minimap2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minimap2_zip_url := "https://github.com/lh3/minimap2/archive/refs/tags/v2.28.zip"
	minimap2_cmd_zip := exec.Command("curl", "-L", minimap2_zip_url, "-o", "package.zip")
	err = minimap2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minimap2_bin_url := "https://github.com/lh3/minimap2/archive/refs/tags/v2.28.bin"
	minimap2_cmd_bin := exec.Command("curl", "-L", minimap2_bin_url, "-o", "binary.bin")
	err = minimap2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minimap2_src_url := "https://github.com/lh3/minimap2/archive/refs/tags/v2.28.src.tar.gz"
	minimap2_cmd_src := exec.Command("curl", "-L", minimap2_src_url, "-o", "source.tar.gz")
	err = minimap2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minimap2_cmd_direct := exec.Command("./binary")
	err = minimap2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
