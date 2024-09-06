package main

// Sratoolkit - Data tools for INSDC Sequence Read Archive
// Homepage: https://github.com/ncbi/sra-tools

import (
	"fmt"
	
	"os/exec"
)

func installSratoolkit() {
	// Método 1: Descargar y extraer .tar.gz
	sratoolkit_tar_url := "https://github.com/ncbi/sra-tools/archive/refs/tags/3.1.1.tar.gz"
	sratoolkit_cmd_tar := exec.Command("curl", "-L", sratoolkit_tar_url, "-o", "package.tar.gz")
	err := sratoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sratoolkit_zip_url := "https://github.com/ncbi/sra-tools/archive/refs/tags/3.1.1.zip"
	sratoolkit_cmd_zip := exec.Command("curl", "-L", sratoolkit_zip_url, "-o", "package.zip")
	err = sratoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sratoolkit_bin_url := "https://github.com/ncbi/sra-tools/archive/refs/tags/3.1.1.bin"
	sratoolkit_cmd_bin := exec.Command("curl", "-L", sratoolkit_bin_url, "-o", "binary.bin")
	err = sratoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sratoolkit_src_url := "https://github.com/ncbi/sra-tools/archive/refs/tags/3.1.1.src.tar.gz"
	sratoolkit_cmd_src := exec.Command("curl", "-L", sratoolkit_src_url, "-o", "source.tar.gz")
	err = sratoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sratoolkit_cmd_direct := exec.Command("./binary")
	err = sratoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
}
