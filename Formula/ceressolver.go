package main

// CeresSolver - C++ library for large-scale optimization
// Homepage: http://ceres-solver.org/

import (
	"fmt"
	
	"os/exec"
)

func installCeresSolver() {
	// Método 1: Descargar y extraer .tar.gz
	ceressolver_tar_url := "http://ceres-solver.org/ceres-solver-2.2.0.tar.gz"
	ceressolver_cmd_tar := exec.Command("curl", "-L", ceressolver_tar_url, "-o", "package.tar.gz")
	err := ceressolver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ceressolver_zip_url := "http://ceres-solver.org/ceres-solver-2.2.0.zip"
	ceressolver_cmd_zip := exec.Command("curl", "-L", ceressolver_zip_url, "-o", "package.zip")
	err = ceressolver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ceressolver_bin_url := "http://ceres-solver.org/ceres-solver-2.2.0.bin"
	ceressolver_cmd_bin := exec.Command("curl", "-L", ceressolver_bin_url, "-o", "binary.bin")
	err = ceressolver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ceressolver_src_url := "http://ceres-solver.org/ceres-solver-2.2.0.src.tar.gz"
	ceressolver_cmd_src := exec.Command("curl", "-L", ceressolver_src_url, "-o", "source.tar.gz")
	err = ceressolver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ceressolver_cmd_direct := exec.Command("./binary")
	err = ceressolver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: metis")
	exec.Command("latte", "install", "metis").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
}
