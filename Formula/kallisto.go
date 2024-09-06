package main

// Kallisto - Quantify abundances of transcripts from RNA-Seq data
// Homepage: https://pachterlab.github.io/kallisto/

import (
	"fmt"
	
	"os/exec"
)

func installKallisto() {
	// Método 1: Descargar y extraer .tar.gz
	kallisto_tar_url := "https://github.com/pachterlab/kallisto/archive/refs/tags/V0.51.0.tar.gz"
	kallisto_cmd_tar := exec.Command("curl", "-L", kallisto_tar_url, "-o", "package.tar.gz")
	err := kallisto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kallisto_zip_url := "https://github.com/pachterlab/kallisto/archive/refs/tags/V0.51.0.zip"
	kallisto_cmd_zip := exec.Command("curl", "-L", kallisto_zip_url, "-o", "package.zip")
	err = kallisto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kallisto_bin_url := "https://github.com/pachterlab/kallisto/archive/refs/tags/V0.51.0.bin"
	kallisto_cmd_bin := exec.Command("curl", "-L", kallisto_bin_url, "-o", "binary.bin")
	err = kallisto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kallisto_src_url := "https://github.com/pachterlab/kallisto/archive/refs/tags/V0.51.0.src.tar.gz"
	kallisto_cmd_src := exec.Command("curl", "-L", kallisto_src_url, "-o", "source.tar.gz")
	err = kallisto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kallisto_cmd_direct := exec.Command("./binary")
	err = kallisto_cmd_direct.Run()
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
