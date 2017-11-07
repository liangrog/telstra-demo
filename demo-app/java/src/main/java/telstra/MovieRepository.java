package telstra;

import org.springframework.data.repository.CrudRepository;

import telstra.Movie;

public interface MovieRepository extends CrudRepository<Movie, Long> {

}
