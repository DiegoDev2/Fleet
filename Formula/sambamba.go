package main

// Sambamba - Tools for working with SAM/BAM data
// Homepage: https://lomereiter.github.io/sambamba/

import (
	"fmt"
	
	"os/exec"
)

func installSambamba() {
	// Método 1: Descargar y extraer .tar.gz
	sambamba_tar_url := "https://github.com/biod/sambamba/archive/refs/tags/v1.0.1.tar.gz"
	sambamba_cmd_tar := exec.Command("curl", "-L", sambamba_tar_url, "-o", "package.tar.gz")
	err := sambamba_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sambamba_zip_url := "https://github.com/biod/sambamba/archive/refs/tags/v1.0.1.zip"
	sambamba_cmd_zip := exec.Command("curl", "-L", sambamba_zip_url, "-o", "package.zip")
	err = sambamba_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sambamba_bin_url := "https://github.com/biod/sambamba/archive/refs/tags/v1.0.1.bin"
	sambamba_cmd_bin := exec.Command("curl", "-L", sambamba_bin_url, "-o", "binary.bin")
	err = sambamba_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sambamba_src_url := "https://github.com/biod/sambamba/archive/refs/tags/v1.0.1.src.tar.gz"
	sambamba_cmd_src := exec.Command("curl", "-L", sambamba_src_url, "-o", "source.tar.gz")
	err = sambamba_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sambamba_cmd_direct := exec.Command("./binary")
	err = sambamba_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ldc")
	exec.Command("latte", "install", "ldc").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
}
