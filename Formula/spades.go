package main

// Spades - De novo genome sequence assembly
// Homepage: https://github.com/ablab/spades

import (
	"fmt"
	
	"os/exec"
)

func installSpades() {
	// Método 1: Descargar y extraer .tar.gz
	spades_tar_url := "https://github.com/ablab/spades/releases/download/v4.0.0/SPAdes-4.0.0.tar.gz"
	spades_cmd_tar := exec.Command("curl", "-L", spades_tar_url, "-o", "package.tar.gz")
	err := spades_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spades_zip_url := "https://github.com/ablab/spades/releases/download/v4.0.0/SPAdes-4.0.0.zip"
	spades_cmd_zip := exec.Command("curl", "-L", spades_zip_url, "-o", "package.zip")
	err = spades_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spades_bin_url := "https://github.com/ablab/spades/releases/download/v4.0.0/SPAdes-4.0.0.bin"
	spades_cmd_bin := exec.Command("curl", "-L", spades_bin_url, "-o", "binary.bin")
	err = spades_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spades_src_url := "https://github.com/ablab/spades/releases/download/v4.0.0/SPAdes-4.0.0.src.tar.gz"
	spades_cmd_src := exec.Command("curl", "-L", spades_src_url, "-o", "source.tar.gz")
	err = spades_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spades_cmd_direct := exec.Command("./binary")
	err = spades_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
	fmt.Println("Instalando dependencia: jemalloc")
	exec.Command("latte", "install", "jemalloc").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
