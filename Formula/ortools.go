package main

// OrTools - Google's Operations Research tools
// Homepage: https://developers.google.com/optimization/

import (
	"fmt"
	
	"os/exec"
)

func installOrTools() {
	// Método 1: Descargar y extraer .tar.gz
	ortools_tar_url := "https://github.com/google/or-tools/archive/refs/tags/v9.10.tar.gz"
	ortools_cmd_tar := exec.Command("curl", "-L", ortools_tar_url, "-o", "package.tar.gz")
	err := ortools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ortools_zip_url := "https://github.com/google/or-tools/archive/refs/tags/v9.10.zip"
	ortools_cmd_zip := exec.Command("curl", "-L", ortools_zip_url, "-o", "package.zip")
	err = ortools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ortools_bin_url := "https://github.com/google/or-tools/archive/refs/tags/v9.10.bin"
	ortools_cmd_bin := exec.Command("curl", "-L", ortools_bin_url, "-o", "binary.bin")
	err = ortools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ortools_src_url := "https://github.com/google/or-tools/archive/refs/tags/v9.10.src.tar.gz"
	ortools_cmd_src := exec.Command("curl", "-L", ortools_src_url, "-o", "source.tar.gz")
	err = ortools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ortools_cmd_direct := exec.Command("./binary")
	err = ortools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: cbc")
	exec.Command("latte", "install", "cbc").Run()
	fmt.Println("Instalando dependencia: cgl")
	exec.Command("latte", "install", "cgl").Run()
	fmt.Println("Instalando dependencia: clp")
	exec.Command("latte", "install", "clp").Run()
	fmt.Println("Instalando dependencia: coinutils")
	exec.Command("latte", "install", "coinutils").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: osi")
	exec.Command("latte", "install", "osi").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
	fmt.Println("Instalando dependencia: scip")
	exec.Command("latte", "install", "scip").Run()
}
