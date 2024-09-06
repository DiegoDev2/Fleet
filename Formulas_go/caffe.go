package main

// Caffe - Fast open framework for deep learning
// Homepage: https://caffe.berkeleyvision.org/

import (
	"fmt"
	
	"os/exec"
)

func installCaffe() {
	// Método 1: Descargar y extraer .tar.gz
	caffe_tar_url := "https://github.com/BVLC/caffe/archive/refs/tags/1.0.tar.gz"
	caffe_cmd_tar := exec.Command("curl", "-L", caffe_tar_url, "-o", "package.tar.gz")
	err := caffe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	caffe_zip_url := "https://github.com/BVLC/caffe/archive/refs/tags/1.0.zip"
	caffe_cmd_zip := exec.Command("curl", "-L", caffe_zip_url, "-o", "package.zip")
	err = caffe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	caffe_bin_url := "https://github.com/BVLC/caffe/archive/refs/tags/1.0.bin"
	caffe_cmd_bin := exec.Command("curl", "-L", caffe_bin_url, "-o", "binary.bin")
	err = caffe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	caffe_src_url := "https://github.com/BVLC/caffe/archive/refs/tags/1.0.src.tar.gz"
	caffe_cmd_src := exec.Command("curl", "-L", caffe_src_url, "-o", "source.tar.gz")
	err = caffe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	caffe_cmd_direct := exec.Command("./binary")
	err = caffe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: leveldb")
exec.Command("latte", "install", "leveldb")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
	fmt.Println("Instalando dependencia: lmdb")
exec.Command("latte", "install", "lmdb")
	fmt.Println("Instalando dependencia: opencv")
exec.Command("latte", "install", "opencv")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
}
