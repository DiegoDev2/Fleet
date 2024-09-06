package main

// Augustus - Predict genes in eukaryotic genomic sequences
// Homepage: https://bioinf.uni-greifswald.de/augustus/

import (
	"fmt"
	
	"os/exec"
)

func installAugustus() {
	// Método 1: Descargar y extraer .tar.gz
	augustus_tar_url := "https://github.com/Gaius-Augustus/Augustus/archive/refs/tags/v3.5.0.tar.gz"
	augustus_cmd_tar := exec.Command("curl", "-L", augustus_tar_url, "-o", "package.tar.gz")
	err := augustus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	augustus_zip_url := "https://github.com/Gaius-Augustus/Augustus/archive/refs/tags/v3.5.0.zip"
	augustus_cmd_zip := exec.Command("curl", "-L", augustus_zip_url, "-o", "package.zip")
	err = augustus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	augustus_bin_url := "https://github.com/Gaius-Augustus/Augustus/archive/refs/tags/v3.5.0.bin"
	augustus_cmd_bin := exec.Command("curl", "-L", augustus_bin_url, "-o", "binary.bin")
	err = augustus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	augustus_src_url := "https://github.com/Gaius-Augustus/Augustus/archive/refs/tags/v3.5.0.src.tar.gz"
	augustus_cmd_src := exec.Command("curl", "-L", augustus_src_url, "-o", "source.tar.gz")
	err = augustus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	augustus_cmd_direct := exec.Command("./binary")
	err = augustus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bamtools")
exec.Command("latte", "install", "bamtools")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: htslib")
exec.Command("latte", "install", "htslib")
}
