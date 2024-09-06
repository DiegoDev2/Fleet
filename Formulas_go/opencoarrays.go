package main

// Opencoarrays - Open-source coarray Fortran ABI, API, and compiler wrapper
// Homepage: http://www.opencoarrays.org

import (
	"fmt"
	
	"os/exec"
)

func installOpencoarrays() {
	// Método 1: Descargar y extraer .tar.gz
	opencoarrays_tar_url := "https://github.com/sourceryinstitute/OpenCoarrays/releases/download/2.10.2/OpenCoarrays-2.10.2.tar.gz"
	opencoarrays_cmd_tar := exec.Command("curl", "-L", opencoarrays_tar_url, "-o", "package.tar.gz")
	err := opencoarrays_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opencoarrays_zip_url := "https://github.com/sourceryinstitute/OpenCoarrays/releases/download/2.10.2/OpenCoarrays-2.10.2.zip"
	opencoarrays_cmd_zip := exec.Command("curl", "-L", opencoarrays_zip_url, "-o", "package.zip")
	err = opencoarrays_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opencoarrays_bin_url := "https://github.com/sourceryinstitute/OpenCoarrays/releases/download/2.10.2/OpenCoarrays-2.10.2.bin"
	opencoarrays_cmd_bin := exec.Command("curl", "-L", opencoarrays_bin_url, "-o", "binary.bin")
	err = opencoarrays_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opencoarrays_src_url := "https://github.com/sourceryinstitute/OpenCoarrays/releases/download/2.10.2/OpenCoarrays-2.10.2.src.tar.gz"
	opencoarrays_cmd_src := exec.Command("curl", "-L", opencoarrays_src_url, "-o", "source.tar.gz")
	err = opencoarrays_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opencoarrays_cmd_direct := exec.Command("./binary")
	err = opencoarrays_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
}
