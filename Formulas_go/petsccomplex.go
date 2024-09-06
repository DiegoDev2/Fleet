package main

// PetscComplex - Portable, Extensible Toolkit for Scientific Computation (complex)
// Homepage: https://petsc.org/

import (
	"fmt"
	
	"os/exec"
)

func installPetscComplex() {
	// Método 1: Descargar y extraer .tar.gz
	petsccomplex_tar_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.tar.gz"
	petsccomplex_cmd_tar := exec.Command("curl", "-L", petsccomplex_tar_url, "-o", "package.tar.gz")
	err := petsccomplex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	petsccomplex_zip_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.zip"
	petsccomplex_cmd_zip := exec.Command("curl", "-L", petsccomplex_zip_url, "-o", "package.zip")
	err = petsccomplex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	petsccomplex_bin_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.bin"
	petsccomplex_cmd_bin := exec.Command("curl", "-L", petsccomplex_bin_url, "-o", "binary.bin")
	err = petsccomplex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	petsccomplex_src_url := "https://web.cels.anl.gov/projects/petsc/download/release-snapshots/petsc-3.20.5.src.tar.gz"
	petsccomplex_cmd_src := exec.Command("curl", "-L", petsccomplex_src_url, "-o", "source.tar.gz")
	err = petsccomplex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	petsccomplex_cmd_direct := exec.Command("./binary")
	err = petsccomplex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hdf5-mpi")
exec.Command("latte", "install", "hdf5-mpi")
	fmt.Println("Instalando dependencia: hwloc")
exec.Command("latte", "install", "hwloc")
	fmt.Println("Instalando dependencia: metis")
exec.Command("latte", "install", "metis")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: scalapack")
exec.Command("latte", "install", "scalapack")
	fmt.Println("Instalando dependencia: suite-sparse")
exec.Command("latte", "install", "suite-sparse")
}
