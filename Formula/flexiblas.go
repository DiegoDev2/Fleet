package main

// Flexiblas - BLAS and LAPACK wrapper library with runtime exchangable backends
// Homepage: https://www.mpi-magdeburg.mpg.de/projects/flexiblas

import (
	"fmt"
	
	"os/exec"
)

func installFlexiblas() {
	// Método 1: Descargar y extraer .tar.gz
	flexiblas_tar_url := "https://csc.mpi-magdeburg.mpg.de/mpcsc/software/flexiblas/flexiblas-3.4.4.tar.xz"
	flexiblas_cmd_tar := exec.Command("curl", "-L", flexiblas_tar_url, "-o", "package.tar.gz")
	err := flexiblas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flexiblas_zip_url := "https://csc.mpi-magdeburg.mpg.de/mpcsc/software/flexiblas/flexiblas-3.4.4.tar.xz"
	flexiblas_cmd_zip := exec.Command("curl", "-L", flexiblas_zip_url, "-o", "package.zip")
	err = flexiblas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flexiblas_bin_url := "https://csc.mpi-magdeburg.mpg.de/mpcsc/software/flexiblas/flexiblas-3.4.4.tar.xz"
	flexiblas_cmd_bin := exec.Command("curl", "-L", flexiblas_bin_url, "-o", "binary.bin")
	err = flexiblas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flexiblas_src_url := "https://csc.mpi-magdeburg.mpg.de/mpcsc/software/flexiblas/flexiblas-3.4.4.tar.xz"
	flexiblas_cmd_src := exec.Command("curl", "-L", flexiblas_src_url, "-o", "source.tar.gz")
	err = flexiblas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flexiblas_cmd_direct := exec.Command("./binary")
	err = flexiblas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
