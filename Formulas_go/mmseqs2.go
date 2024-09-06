package main

// Mmseqs2 - Software suite for very fast sequence search and clustering
// Homepage: https://mmseqs.com/

import (
	"fmt"
	
	"os/exec"
)

func installMmseqs2() {
	// Método 1: Descargar y extraer .tar.gz
	mmseqs2_tar_url := "https://github.com/soedinglab/MMseqs2/archive/refs/tags/15-6f452.tar.gz"
	mmseqs2_cmd_tar := exec.Command("curl", "-L", mmseqs2_tar_url, "-o", "package.tar.gz")
	err := mmseqs2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmseqs2_zip_url := "https://github.com/soedinglab/MMseqs2/archive/refs/tags/15-6f452.zip"
	mmseqs2_cmd_zip := exec.Command("curl", "-L", mmseqs2_zip_url, "-o", "package.zip")
	err = mmseqs2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmseqs2_bin_url := "https://github.com/soedinglab/MMseqs2/archive/refs/tags/15-6f452.bin"
	mmseqs2_cmd_bin := exec.Command("curl", "-L", mmseqs2_bin_url, "-o", "binary.bin")
	err = mmseqs2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmseqs2_src_url := "https://github.com/soedinglab/MMseqs2/archive/refs/tags/15-6f452.src.tar.gz"
	mmseqs2_cmd_src := exec.Command("curl", "-L", mmseqs2_src_url, "-o", "source.tar.gz")
	err = mmseqs2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmseqs2_cmd_direct := exec.Command("./binary")
	err = mmseqs2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
}
